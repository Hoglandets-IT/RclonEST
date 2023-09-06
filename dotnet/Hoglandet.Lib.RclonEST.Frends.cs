using System.Diagnostics;
using System.ComponentModel;
using System.Text.RegularExpressions;
using System.Text;
using Newtonsoft.Json;
using System.Threading.Tasks;
using Hoglandet.Lib.RclonEST;

/// <summary>
/// Make requests to a HTTP-enabled Rclone instance
/// </summary>
namespace Hoglandet.Lib.RclonEST
{
    /// <summary>
    /// Action class for running commands against Rclone
    /// </summary>
    public static class Frends
    {
        /// <summary>
        /// Ensure the configuration exists on the rclone side and is up to date
        /// </summary>
        /// <param name="rem"></param>
        /// <param name="conn"></param>
        /// <returns></returns>
        public static bool EnsureConfigurationExists([PropertyTab]Remote rem, [PropertyTab]ConnectionSettings conn) 
        {
            var client = new RclonEST.Client(conn.Host, conn.User, conn.Password);
            rem.Upsert(client);

            return true;
        }

        /// <summary>
        /// List all files on the given path on the remote
        /// </summary>
        /// <param name="settings"></param>
        /// <param name="remote"></param>
        /// <param name="conn"></param>
        /// <returns></returns>
        public static dynamic ListFiles([PropertyTab]ListSettings settings, [PropertyTab]RemoteSettings remote, [PropertyTab]ConnectionSettings conn) {
            var cli = new RclonEST.Client(conn.Host, conn.User, conn.Password);

            Remote remoteCfg = remote.GetRemote();

            var cmd = new RclonEST.OperationsList(){
                Remote = remote.Name.ToLower() + ":",
                Path = settings.Path
            };

            var res = cli.RunCommand(cmd);
            return res;
        }

        public static string ListRemotes([PropertyTab]ConnectionSettings conn) {
            var cli = new RclonEST.Client(conn.Host, conn.User, conn.Password);
            var cmd = new RclonEST.ConfigListremotes();
            var res = cli.RunCommand(cmd);
            return "res";
        }
    }
}
