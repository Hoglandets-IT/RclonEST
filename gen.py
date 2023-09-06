import os
import yaml
import regex as re

# with open('/home/larsch/rclone-caller/test-genfiles/config.yaml', 'r') as f:
with open('api-struct.yaml', 'r') as f:
    config_all = yaml.safe_load(f)
    config_endpoints = config_all['endpoints']

if not os.path.isdir('golang/endpoints'):
    os.mkdir('golang/endpoints')

struct_template = '''
// {description}
type {structname} struct {{
    JsonData string `json:"-"`
{structdata} 
}}
'''
'''
//export Run{structname}
func Run{structname}(credentials string, {structdata_call}) (*C.char, *C.char) {{
    var cred RCCredentials
	Unmarshal(credentials, &cred)
	var call {structname} = {structname}{{
        {structdata_assign}
    }}
}}

'''


net_file = '''
using Newtonsoft.Json;

namespace Hoglandet.Lib.RclonEST {
public class CallBase {
    public string GetPath() {
        return "/" + System.Text.RegularExpressions.Regex.Replace(this.GetType().Name, "(?<=.)([A-Z])", "/$0",
                    System.Text.RegularExpressions.RegexOptions.Compiled).ToLowerInvariant();
    }
    
    public string GetJson() {
        return JsonConvert.SerializeObject(this);
    }
}

'''


net_struct_template = '''
public class {structname} : CallBase {{
    [JsonIgnore]
    public string JsonData {{ get; set; }} = "";
{structdatanet}
}}

'''

net_remotes_file = '''
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Data.Common;
using System.Reflection;
using System.Text;
using Newtonsoft.Json;

namespace Hoglandet.Lib.RclonEST {{
    public class RemoteOpts {{
        [JsonProperty("obscure", NullValueHandling=NullValueHandling.Ignore)]
        public bool? Obscure {{get; set;}} = null;
        
        [JsonProperty("noObscure", NullValueHandling=NullValueHandling.Ignore)]
        public bool? NoObscure {{get; set;}} = null;
        
        [JsonProperty("nonInteractive", NullValueHandling=NullValueHandling.Ignore)]
        public bool? NonInteractive {{get; set;}} = null;
        
        [JsonProperty("continue", NullValueHandling=NullValueHandling.Ignore)]
        public bool? Continue {{get; set;}} = null;
        
        [JsonProperty("all", NullValueHandling=NullValueHandling.Ignore)]
        public bool? All {{get; set;}} = null;
        
        [JsonProperty("state", NullValueHandling=NullValueHandling.Ignore)]
        public bool? State {{get; set;}} = null;
        
        [JsonProperty("result", NullValueHandling=NullValueHandling.Ignore)]
        public bool? Result {{get; set;}} = null;
        
        public Dictionary<string, bool> GetDict() {{
			return JsonConvert.DeserializeObject<Dictionary<string, bool>>(
				JsonConvert.SerializeObject(this)
			);
		}}
    }}
    
    public enum RemoteType {{
        {remotes}
    }}
    
    
    public class Remote {{
        [JsonProperty("name")]
        public string Name {{ get; set; }}
        
        [JsonProperty("type")]
        public RemoteType Type {{ get; set; }}
        
        [JsonProperty("parameters", NullValueHandling=NullValueHandling.Ignore)]
        public Dictionary<string, string> Parameters {{ get; set; }}
        
        [JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
        public RemoteOpts Opt {{ get; set; }}
        
        public void Create(RclonEST.Client client) {{
			var cfg = new ConfigCreate() {{
					Name = Name.ToLower(),
					Type = Type.ToString().ToLower(),
					Parameters = Parameters,
			}};

			if (Opt.GetDict().Count < 0) {{
				cfg.Options = Opt.GetDict();
			}}

			var cmd = client.RunCommand(cfg);
			if (cmd.error != null && cmd.error != "") {{
				Console.WriteLine("ERROR: ", cmd.error);
			}}
		}}

		public bool Exists(RclonEST.Client client) {{
			return client.RunCommand(
				new ConfigGet() {{
					Name = Name.ToLower()
				}}
			).Count > 0;
		}}

		public void Update(RclonEST.Client client) {{
			var cfg = new ConfigUpdate() {{
				Name = Name.ToLower(),
				Parameters = Parameters
			}};
			
			if (Opt.GetDict().Count < 0) {{
				cfg.Options = Opt.GetDict();
			}}

		}}

		public void Upsert(RclonEST.Client client) {{
			if (!this.Exists(client)) {{
				this.Create(client);
				return;
			}}
			this.Update(client);
			return;
		}}
    }}
}}
'''

net_remotes_template = '''
    \t[Display(Name = "{remotenamefmt}")]
    \t{remotename},'''

# net_remotes_template = '''
#     public class {remotename} : RemoteBase {{
#         [JsonProperty("type")]
#         public string Type {{
#             get {{
#                 return "{remotetype}";
#             }}
#         }}
#     }}

# '''


capi_file = '''
package main

import "C"

import (
	"encoding/json"
	"github.com/hoglandets-it/rclonest/golang/endpoints"
)

func Unmarshal(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}

//func RunGeneric(credentials string, jsondata string, endpoint interface{}) (string, string) {
//    Unmarshal(jsondata, &endpoint)
    
//    cl := ClientFromJson(credentials)
//    res, err := cl.Run(endpoint)
//    if (err != nil) {
//        return "", string(err.Error())
//    }
//    return string(res), ""
//}


func RunGeneric(credentials string, jsondata string, path string) (string, string) {
	var creds map[string]string
	json.Unmarshal([]byte(credentials), &creds)

    //jsVal := Remarshal(jsondata, &endpoint)


	// // x := reflect.ValueOf(&endpoint)
	// // fmt.Println(x)

	if creds["url"] == "" || creds["username"] == "" || creds["password"] == "" {
		return "", "Missing credentials"
	}

	client := NewClient(creds["url"], creds["username"], creds["password"])
	res, err := client.RunC()
	// if err != nil {
	// 	return "", string(err.Error())
	// }
	// return string(res), ""

}

'''
capi_template = '''
//export Run{structname}
func Run{structname} (credentials string, json_data string) (string, string) {{
    var cm{structname} endpoints.{structname}
    json.Unmarshal([]byte(json_data), &cm{structname})
    
    return RunGeneric(credentials, json_data, "{pathname}")
}}

'''

for namespace in config_endpoints:
    structs = []
    netstructs = []
    for funcKey, funcVal in config_endpoints[namespace].items():
        template = str(struct_template)
        funcDef = {
            "description": funcVal['description'],
            "structname": f"{namespace.title()}{funcKey.title()}",
            "pathname": f"{namespace.lower()}/{funcKey.lower()}",
            "structdata": "",
            "structdatanet": ""
        }
        
        if '-' in funcDef['structname']:
            funcDef['structname'] = funcDef['structname'].replace('-', '')
        
        for argKey, argVal in funcVal['args'].items():
            argName = argVal['name']
            argType = argVal['type'].split(',')
            argTypeNet = argType
            argTypeCall = argVal['type'].split(',')
            
            if (argType[0] == 'string' or argType[0] == 'bool' or argType[0] == 'int'):
                argTypeCall = argType[0]
            else:
                argTypeCall = 'string'
            
            if len(argType) == 1 and argType[0]:
                argType = argType[0]
                argTypeNet = argTypeNet[0]
            elif len(argType) == 2 and argType[0] in ['list', 'array', 'set']:
                argType = f"[]{argType[1]}"
                argTypeNet = f"{argTypeNet[1]}[]"
            elif len(argType) == 3:
                argType = f"map[{argType[1]}]{argType[2]}"
                argTypeNet = f"Dictionary<{argTypeNet[1]}, {argTypeNet[2]}>"
            else:
                raise Exception(f"Invalid type {argType}")
            
            if argType == 'json':
                argType = 'string'
                argTypeNet = 'string'
            
            funcDef["structdata"] += f"    {argVal['name']} {argType} `json:\"{argKey},omitempty\"`\n"
            funcDef["structdatanet"] += f"[JsonProperty(\"{argKey}\", NullValueHandling=NullValueHandling.Ignore)]\n    public {argTypeNet} {argVal['name']} {{ get; set; }}\n"
        
        capi_file += capi_template.format(**funcDef)
        net_file += net_struct_template.format(**funcDef)
        structs.append(struct_template.format(**funcDef))
    
    with open('golang/endpoints/' + namespace + '.go', 'w') as f:
        f.write("package endpoints \n\n")
        f.write(
            '\n\n'.join(structs)
        )
    
    with open('dotnet/Hoglandet.Lib.RclonEST.Definitions.cs', 'w') as f:
        f.write(
            net_file + "}"
        )
    
    with open('golang/capi.go', 'w') as f:
        f.write(capi_file)
    
net_remotes = ""

for remoteKey in config_all['remoteTypes']:
    net_remotes += "\n" + net_remotes_template.format(
        remotename=remoteKey.upper(),
        remotenamefmt=re.sub(r'(\w)([A-Z])', r'\1 \2', remoteKey).title()
    )
    
with open('dotnet/Hoglandet.Lib.RclonEST.Remotes.cs', 'w') as f:
    f.write(
        net_remotes_file.format(remotes=net_remotes)
    )
    

