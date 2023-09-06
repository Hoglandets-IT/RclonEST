package endpoints 


// Run a bilateral sync between two paths/remotes/filesystems
type SyncBisync struct {
    JsonData string `json:"-"`
    PathOne string `json:"path1,omitempty"`
    PathTwo string `json:"path2,omitempty"`
    DryRun bool `json:"dryRun,omitempty"`
    Resync bool `json:"resync,omitempty"`
    CheckAccess bool `json:"checkAccess,omitempty"`
    MaxDelete int `json:"maxDelete,omitempty"`
    Force bool `json:"force,omitempty"`
    CheckSync string `json:"checkSync,omitempty"`
    RemoveEmptyDirs bool `json:"removeEmptyDirs,omitempty"`
    FiltersFile string `json:"filtersFile,omitempty"`
    WorkDir string `json:"workdir,omitempty"`
    NoCleanup bool `json:"noCleanup,omitempty"`
 
}



// Copy a directory between source remote and destination remote
type SyncCopy struct {
    JsonData string `json:"-"`
    SourceRemoteString string `json:"srcFs,omitempty"`
    DestinationRemoteString string `json:"dstFs,omitempty"`
    CreateEmptySrcDirs bool `json:"createEmptySrcDirs,omitempty"`
 
}



// Move a directory from soruce remote to destination remote
type SyncMove struct {
    JsonData string `json:"-"`
    SourceRemoteString string `json:"srcFs,omitempty"`
    DestinationRemoteString string `json:"dstFs,omitempty"`
    CreateEmptySrcDirs bool `json:"createEmptySrcDirs,omitempty"`
    DeleteEmptySrcDirs bool `json:"deleteEmptySrcDirs,omitempty"`
 
}



// Sync a directory between source and destination filesystem
type SyncSync struct {
    JsonData string `json:"-"`
    SourceRemoteString string `json:"srcFs,omitempty"`
    DestinationRemoteString string `json:"dstFs,omitempty"`
    CreateEmptySrcDirs bool `json:"createEmptySrcDirs,omitempty"`
 
}
