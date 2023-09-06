using RestSharp;
using RestSharp.Authenticators;
using Newtonsoft.Json;

namespace Hoglandet.Lib.RclonEST
{
    public class Client
    {
        private static RestClientOptions rco;
        private static RestClient rc;

        public Client(string host, string username, string password) {
            string fmtHost = host;
            
            if (!fmtHost.StartsWith('https')) {
                fmtHost = "https://" + fmtHost;
            }


            rco = new RestClientOptions(host) {
                Authenticator = new HttpBasicAuthenticator(username, password)
            };

            rc = new RestClient(rco);
        }
        
        public dynamic RunCommand(CallBase argm)
        {
            var request = new RestRequest(argm.GetPath(), Method.Post).AddBody(
                argm.GetJson()
            );

            request.AddHeader("Content-Type", "application/json");
            var resp = rc.Execute(request);

            return JsonConvert.DeserializeObject(resp.Content);
        }
    }
}
