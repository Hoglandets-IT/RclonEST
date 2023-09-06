
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Data.Common;
using System.Reflection;
using System.Text;
using Newtonsoft.Json;

namespace Hoglandet.Lib.RclonEST {
    public class RemoteOpts {
        [JsonProperty("obscure", NullValueHandling=NullValueHandling.Ignore)]
        public bool? Obscure {get; set;} = null;
        
        [JsonProperty("noObscure", NullValueHandling=NullValueHandling.Ignore)]
        public bool? NoObscure {get; set;} = null;
        
        [JsonProperty("nonInteractive", NullValueHandling=NullValueHandling.Ignore)]
        public bool? NonInteractive {get; set;} = null;
        
        [JsonProperty("continue", NullValueHandling=NullValueHandling.Ignore)]
        public bool? Continue {get; set;} = null;
        
        [JsonProperty("all", NullValueHandling=NullValueHandling.Ignore)]
        public bool? All {get; set;} = null;
        
        [JsonProperty("state", NullValueHandling=NullValueHandling.Ignore)]
        public bool? State {get; set;} = null;
        
        [JsonProperty("result", NullValueHandling=NullValueHandling.Ignore)]
        public bool? Result {get; set;} = null;
        
        public Dictionary<string, bool> GetDict() {
			return JsonConvert.DeserializeObject<Dictionary<string, bool>>(
				JsonConvert.SerializeObject(this)
			);
		}
    }
    
    public enum RemoteType {
        

    	[Display(Name = "Amazon Clouddrive")]
    	AMAZONCLOUDDRIVE,

    	[Display(Name = "Azure Blob")]
    	AZUREBLOB,

    	[Display(Name = "B2")]
    	B2,

    	[Display(Name = "Box")]
    	BOX,

    	[Display(Name = "Cache")]
    	CACHE,

    	[Display(Name = "Chunker")]
    	CHUNKER,

    	[Display(Name = "Combine")]
    	COMBINE,

    	[Display(Name = "Compress")]
    	COMPRESS,

    	[Display(Name = "Crypt")]
    	CRYPT,

    	[Display(Name = "Drive")]
    	DRIVE,

    	[Display(Name = "Dropbox")]
    	DROPBOX,

    	[Display(Name = "Fichier")]
    	FICHIER,

    	[Display(Name = "Filefabric")]
    	FILEFABRIC,

    	[Display(Name = "Ftp")]
    	FTP,

    	[Display(Name = "Google Cloud Storage")]
    	GOOGLECLOUDSTORAGE,

    	[Display(Name = "Google Photos")]
    	GOOGLEPHOTOS,

    	[Display(Name = "Hasher")]
    	HASHER,

    	[Display(Name = "Hdfs")]
    	HDFS,

    	[Display(Name = "Hidrive")]
    	HIDRIVE,

    	[Display(Name = "Http")]
    	HTTP,

    	[Display(Name = "Internet Archive")]
    	INTERNETARCHIVE,

    	[Display(Name = "Jottacloud")]
    	JOTTACLOUD,

    	[Display(Name = "Koofr")]
    	KOOFR,

    	[Display(Name = "Local")]
    	LOCAL,

    	[Display(Name = "Mailru")]
    	MAILRU,

    	[Display(Name = "Mega")]
    	MEGA,

    	[Display(Name = "Memory")]
    	MEMORY,

    	[Display(Name = "Netstorage")]
    	NETSTORAGE,

    	[Display(Name = "One Drive")]
    	ONEDRIVE,

    	[Display(Name = "Open Drive")]
    	OPENDRIVE,

    	[Display(Name = "Oracle Object Storage")]
    	ORACLEOBJECTSTORAGE,

    	[Display(Name = "Pcloud")]
    	PCLOUD,

    	[Display(Name = "Pikpak")]
    	PIKPAK,

    	[Display(Name = "Premium Izeme")]
    	PREMIUMIZEME,

    	[Display(Name = "Protondrive")]
    	PROTONDRIVE,

    	[Display(Name = "Putio")]
    	PUTIO,

    	[Display(Name = "Qingstor")]
    	QINGSTOR,

    	[Display(Name = "S3")]
    	S3,

    	[Display(Name = "Seafile")]
    	SEAFILE,

    	[Display(Name = "Sftp")]
    	SFTP,

    	[Display(Name = "Sharefile")]
    	SHAREFILE,

    	[Display(Name = "Sia")]
    	SIA,

    	[Display(Name = "Smb")]
    	SMB,

    	[Display(Name = "Storj")]
    	STORJ,

    	[Display(Name = "Sugarsync")]
    	SUGARSYNC,

    	[Display(Name = "Swift")]
    	SWIFT,

    	[Display(Name = "Union")]
    	UNION,

    	[Display(Name = "Uptobox")]
    	UPTOBOX,

    	[Display(Name = "Webdav")]
    	WEBDAV,

    	[Display(Name = "Yandex")]
    	YANDEX,

    	[Display(Name = "Zoho")]
    	ZOHO,
    }
    
    
    public class Remote {
        [JsonProperty("name")]
        public string Name { get; set; }
        
        [JsonProperty("type")]
        public RemoteType Type { get; set; }
        
        [JsonProperty("parameters", NullValueHandling=NullValueHandling.Ignore)]
        public Dictionary<string, string> Parameters { get; set; }
        
        [JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
        public RemoteOpts Opt { get; set; }
        
        public void Create(RclonEST.Client client) {
			var cfg = new ConfigCreate() {
					Name = Name.ToLower(),
					Type = Type.ToString().ToLower(),
					Parameters = Parameters,
			};

			if (Opt.GetDict().Count < 0) {
				cfg.Options = Opt.GetDict();
			}

			var cmd = client.RunCommand(cfg);
			if (cmd.error != null && cmd.error != "") {
				Console.WriteLine("ERROR: ", cmd.error);
			}
		}

		public bool Exists(RclonEST.Client client) {
			return client.RunCommand(
				new ConfigGet() {
					Name = Name.ToLower()
				}
			).Count > 0;
		}

		public void Update(RclonEST.Client client) {
			var cfg = new ConfigUpdate() {
				Name = Name.ToLower(),
				Parameters = Parameters
			};
			
			if (Opt.GetDict().Count < 0) {
				cfg.Options = Opt.GetDict();
			}

		}

		public void Upsert(RclonEST.Client client) {
			if (!this.Exists(client)) {
				this.Create(client);
				return;
			}
			this.Update(client);
			return;
		}
    }
}
