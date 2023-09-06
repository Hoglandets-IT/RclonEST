using Microsoft.VisualBasic;
namespace Hoglandet.Lib.RclonEST {
    /// <summary>
    /// Connection settings for connecting to the Rclone instance
    /// </summary>
    public class ConnectionSettings {

        /// <summary>
        /// The address for Rclone service
        /// </summary>
        /// <value></value>
        public string Host { get; set; } = "";

        /// <summary>
        /// The username for the Rclone service
        /// </summary>
        /// <value></value>
        public string User { get; set; } = "";

        /// <summary>
        /// The password for the rclone service
        /// </summary>
        /// <value></value>
        public string Password { get; set; } = "";
    }

    public class RemoteSettings {
        /// <summary>
        /// Type of remote
        /// </summary> 
        public RemoteType Type { get; set; }

        /// <summary>
        /// Get the configuration from an environment variable
        /// </summary>
        /// <value></value>
        [Display(Name = "From Environment")]
        [DefaultValue(false)]
        public bool FromEnvironment { get; set; }

        [Display(nameof = "Environment Variable")]
        [UIHint(nameof(FromEnvironment), "", true)]
        public string? EnvironmentVariable { get; set; }

        /// <summary>
        /// The domain name for the user connecting to the SMB Server
        /// </summary>
        /// <value></value>
        [Display(nameof = "LDAP/AD Domain")]
        [UIHint(nameof(FromEnvironment), "", false)]
        [UIHint(nameof(Type), "", RemoteType.SMB)]
        public string? Domain { get; set; }

        /// <summary>
        /// The address for the remote, e.g. ftp.domain.com or hinternal01.intern.domain.com
        /// </summary>
        [UIHint(nameof(FromEnvironment), "", false)]
        public string? Address { get; set; }

        [UIHint(nameof(FromEnvironment), "", false)]
        public string? Username { get; set; }

        [UIHint(nameof(FromEnvironment), "", false)]
        public string? Password { get; set; }

        [UIHint(nameof(FromEnvironment), "", false)]
        public string? PrivateKey { get; set; }

        [UIHint(nameof(FromEnvironment), "", false)]
        public string? PrivateKeyPassword { get; set; }

        [UIHint(nameof(FromEnvironment), "", false)]
        public string? Fingerprint { get; set; }

        public GetRemote() {
            var re = new Remote();
            var reS = this;

            if (!string.IsNullOrWhiteSpace(reS.EnvironmentVariable) && reS.FromEnvironment) {
                reS = JsonConvert.DeserializeObject<RemoteSettings>(reS.EnvironmentVariable);
            }

            

        }

    }

    public class ListSettings {
        /// <summary>
        /// The path on the remote to list
        /// </summary>
        /// <value></value>
        public string Path { get; set; }

        /// <summary>
        /// Include folders in the listing
        /// </summary>
        /// <value></value>
        public bool IncludeFolders { get; set; } = false;

        /// <summary>
        /// Include files in the listing
        /// </summary>
        /// <value></value>
        public bool IncludeFiles { get; set; } = true;        

    }
}