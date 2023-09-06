
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


public class BackendCommand : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("command", NullValueHandling=NullValueHandling.Ignore)]
    public string Command { get; set; }
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("arg", NullValueHandling=NullValueHandling.Ignore)]
    public string[] Arguments { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> Options { get; set; }

}


public class CacheExpire : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("withData", NullValueHandling=NullValueHandling.Ignore)]
    public bool WithData { get; set; }

}


public class CacheFetch : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class CacheStats : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class ConfigCreate : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string Name { get; set; }
[JsonProperty("parameters", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> Parameters { get; set; }
[JsonProperty("type", NullValueHandling=NullValueHandling.Ignore)]
    public string Type { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, bool> Options { get; set; }

}


public class ConfigDelete : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string Name { get; set; }

}


public class ConfigDump : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class ConfigGet : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string Name { get; set; }

}


public class ConfigListremotes : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class ConfigPassword : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string Name { get; set; }
[JsonProperty("parameters", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> Parameters { get; set; }

}


public class ConfigProviders : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class ConfigSetpath : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("path", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }

}


public class ConfigUpdate : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string Name { get; set; }
[JsonProperty("parameters", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> Parameters { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, bool> Options { get; set; }

}


public class CoreBwlimit : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("rate", NullValueHandling=NullValueHandling.Ignore)]
    public string Rate { get; set; }

}


public class CoreCommand : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("command", NullValueHandling=NullValueHandling.Ignore)]
    public string Command { get; set; }
[JsonProperty("arg", NullValueHandling=NullValueHandling.Ignore)]
    public string[] Arguments { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> Options { get; set; }
[JsonProperty("returnType", NullValueHandling=NullValueHandling.Ignore)]
    public string ReturnType { get; set; }

}


public class CoreGc : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class CoreGroupList : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class CoreMemstats : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class CoreObscure : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("clear", NullValueHandling=NullValueHandling.Ignore)]
    public string Clear { get; set; }

}


public class CorePid : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class CoreQuit : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("exitCode", NullValueHandling=NullValueHandling.Ignore)]
    public int ExitCode { get; set; }

}


public class CoreStats : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("group", NullValueHandling=NullValueHandling.Ignore)]
    public string Group { get; set; }

}


public class CoreStatsDelete : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("group", NullValueHandling=NullValueHandling.Ignore)]
    public string Group { get; set; }

}


public class CoreStatsReset : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("group", NullValueHandling=NullValueHandling.Ignore)]
    public string Group { get; set; }

}


public class CoreTransferred : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("group", NullValueHandling=NullValueHandling.Ignore)]
    public string Group { get; set; }

}


public class CoreVersion : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class DebugSetBlockProfileRate : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("rate", NullValueHandling=NullValueHandling.Ignore)]
    public int Rate { get; set; }

}


public class DebugSetGcPercent : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("rate", NullValueHandling=NullValueHandling.Ignore)]
    public int Rate { get; set; }

}


public class DebugSetMutexProfileFraction : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("rate", NullValueHandling=NullValueHandling.Ignore)]
    public int Rate { get; set; }

}


public class DebugSetSoftMemoryLimit : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("mem-limit", NullValueHandling=NullValueHandling.Ignore)]
    public string MemLimit { get; set; }

}


public class FscacheClear : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class FscacheEntries : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class JobList : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class JobStatus : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("jobid", NullValueHandling=NullValueHandling.Ignore)]
    public int JobID { get; set; }

}


public class JobStop : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("jobid", NullValueHandling=NullValueHandling.Ignore)]
    public int JobID { get; set; }

}


public class JobStopgroup : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("group", NullValueHandling=NullValueHandling.Ignore)]
    public string Group { get; set; }

}


public class MountListmounts : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class MountMount : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("mountPoint", NullValueHandling=NullValueHandling.Ignore)]
    public string MountPoint { get; set; }
[JsonProperty("mountType", NullValueHandling=NullValueHandling.Ignore)]
    public string MountType { get; set; }
[JsonProperty("mountOpt", NullValueHandling=NullValueHandling.Ignore)]
    public string MountOptions { get; set; }
[JsonProperty("vfsOpt", NullValueHandling=NullValueHandling.Ignore)]
    public string VfsOptions { get; set; }

}


public class MountTypes : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class MountUnmount : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("mountPoint", NullValueHandling=NullValueHandling.Ignore)]
    public string MountPoint { get; set; }

}


public class MountUnmountall : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class OperationsAbout : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }

}


public class OperationsCleanup : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }

}


public class OperationsCopyfile : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("srcFs", NullValueHandling=NullValueHandling.Ignore)]
    public string SourceRemote { get; set; }
[JsonProperty("srcRemote", NullValueHandling=NullValueHandling.Ignore)]
    public string SourcePath { get; set; }
[JsonProperty("dstFs", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationRemote { get; set; }
[JsonProperty("dstRemote", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationPath { get; set; }

}


public class OperationsCopyurl : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("url", NullValueHandling=NullValueHandling.Ignore)]
    public string URL { get; set; }
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationRemote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationPath { get; set; }
[JsonProperty("autoFilename", NullValueHandling=NullValueHandling.Ignore)]
    public bool AutomaticFilename { get; set; }

}


public class OperationsDelete : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string RemoteString { get; set; }

}


public class OperationsDeletefile : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }

}


public class OperationsFsinfo : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string RemoteString { get; set; }

}


public class OperationsList : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> ListOptions { get; set; }

}


public class OperationsMkdir : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }

}


public class OperationsMovefile : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("srcFs", NullValueHandling=NullValueHandling.Ignore)]
    public string SourceRemote { get; set; }
[JsonProperty("srcRemote", NullValueHandling=NullValueHandling.Ignore)]
    public string SourcePath { get; set; }
[JsonProperty("dstFs", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationRemote { get; set; }
[JsonProperty("dstRemote", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationPath { get; set; }

}


public class OperationsPubliclink : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }

}


public class OperationsPurge : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }

}


public class OperationsRmdir : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }

}


public class OperationsRmdirs : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }
[JsonProperty("leaveRoot", NullValueHandling=NullValueHandling.Ignore)]
    public bool LeaveRoot { get; set; }

}


public class OperationsSize : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string RemotePath { get; set; }

}


public class OperationsStat : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> StatOptions { get; set; }

}


public class OperationsUploadfile : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("fs", NullValueHandling=NullValueHandling.Ignore)]
    public string Remote { get; set; }
[JsonProperty("remote", NullValueHandling=NullValueHandling.Ignore)]
    public string Path { get; set; }
[JsonProperty("file", NullValueHandling=NullValueHandling.Ignore)]
    public string File { get; set; }
[JsonProperty("opt", NullValueHandling=NullValueHandling.Ignore)]
    public Dictionary<string, string> UploadOptions { get; set; }

}


public class OptionsBlocks : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class OptionsGet : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class OptionsLocal : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class OptionsSet : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class PluginsctlAddplugin : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("url", NullValueHandling=NullValueHandling.Ignore)]
    public string PluginUrl { get; set; }

}


public class PluginsctlGetpluginsfortype : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("type", NullValueHandling=NullValueHandling.Ignore)]
    public string MimeType { get; set; }
[JsonProperty("pluginType", NullValueHandling=NullValueHandling.Ignore)]
    public string PluginType { get; set; }

}


public class PluginsctlListplugins : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class PluginsctlListtestplugins : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class PluginsctlRemoveplugin : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string PluginName { get; set; }

}


public class PluginsctlRemovetestplugin : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("name", NullValueHandling=NullValueHandling.Ignore)]
    public string PluginName { get; set; }

}


public class RcError : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class RcList : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class RcNoop : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class RcNoopauth : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class SyncBisync : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("path1", NullValueHandling=NullValueHandling.Ignore)]
    public string PathOne { get; set; }
[JsonProperty("path2", NullValueHandling=NullValueHandling.Ignore)]
    public string PathTwo { get; set; }
[JsonProperty("dryRun", NullValueHandling=NullValueHandling.Ignore)]
    public bool DryRun { get; set; }
[JsonProperty("resync", NullValueHandling=NullValueHandling.Ignore)]
    public bool Resync { get; set; }
[JsonProperty("checkAccess", NullValueHandling=NullValueHandling.Ignore)]
    public bool CheckAccess { get; set; }
[JsonProperty("maxDelete", NullValueHandling=NullValueHandling.Ignore)]
    public int MaxDelete { get; set; }
[JsonProperty("force", NullValueHandling=NullValueHandling.Ignore)]
    public bool Force { get; set; }
[JsonProperty("checkSync", NullValueHandling=NullValueHandling.Ignore)]
    public string CheckSync { get; set; }
[JsonProperty("removeEmptyDirs", NullValueHandling=NullValueHandling.Ignore)]
    public bool RemoveEmptyDirs { get; set; }
[JsonProperty("filtersFile", NullValueHandling=NullValueHandling.Ignore)]
    public string FiltersFile { get; set; }
[JsonProperty("workdir", NullValueHandling=NullValueHandling.Ignore)]
    public string WorkDir { get; set; }
[JsonProperty("noCleanup", NullValueHandling=NullValueHandling.Ignore)]
    public bool NoCleanup { get; set; }

}


public class SyncCopy : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("srcFs", NullValueHandling=NullValueHandling.Ignore)]
    public string SourceRemoteString { get; set; }
[JsonProperty("dstFs", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationRemoteString { get; set; }
[JsonProperty("createEmptySrcDirs", NullValueHandling=NullValueHandling.Ignore)]
    public bool CreateEmptySrcDirs { get; set; }

}


public class SyncMove : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("srcFs", NullValueHandling=NullValueHandling.Ignore)]
    public string SourceRemoteString { get; set; }
[JsonProperty("dstFs", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationRemoteString { get; set; }
[JsonProperty("createEmptySrcDirs", NullValueHandling=NullValueHandling.Ignore)]
    public bool CreateEmptySrcDirs { get; set; }
[JsonProperty("deleteEmptySrcDirs", NullValueHandling=NullValueHandling.Ignore)]
    public bool DeleteEmptySrcDirs { get; set; }

}


public class SyncSync : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("srcFs", NullValueHandling=NullValueHandling.Ignore)]
    public string SourceRemoteString { get; set; }
[JsonProperty("dstFs", NullValueHandling=NullValueHandling.Ignore)]
    public string DestinationRemoteString { get; set; }
[JsonProperty("createEmptySrcDirs", NullValueHandling=NullValueHandling.Ignore)]
    public bool CreateEmptySrcDirs { get; set; }

}


public class VfsForget : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class VfsList : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class VfsPollInterval : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";
[JsonProperty("interval", NullValueHandling=NullValueHandling.Ignore)]
    public string Interval { get; set; }

}


public class VfsRefresh : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}


public class VfsStats : CallBase {
    [JsonIgnore]
    public string JsonData { get; set; } = "";

}

}