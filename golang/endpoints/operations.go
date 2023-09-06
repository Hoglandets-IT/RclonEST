package endpoints 


// Return the space used on the remote
type OperationsAbout struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
 
}



// Clean up the trashed files in the remote or path
type OperationsCleanup struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
 
}



// Copy a file between filesystems
type OperationsCopyfile struct {
    JsonData string `json:"-"`
    SourceRemote string `json:"srcFs,omitempty"`
    SourcePath string `json:"srcRemote,omitempty"`
    DestinationRemote string `json:"dstFs,omitempty"`
    DestinationPath string `json:"dstRemote,omitempty"`
 
}



// Copy url content to dest
type OperationsCopyurl struct {
    JsonData string `json:"-"`
    URL string `json:"url,omitempty"`
    DestinationRemote string `json:"fs,omitempty"`
    DestinationPath string `json:"remote,omitempty"`
    AutomaticFilename bool `json:"autoFilename,omitempty"`
 
}



// Remove files in a path
type OperationsDelete struct {
    JsonData string `json:"-"`
    RemoteString string `json:"fs,omitempty"`
 
}



// Remove a single file from remote
type OperationsDeletefile struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
 
}



// Print filesystem info for the remote
type OperationsFsinfo struct {
    JsonData string `json:"-"`
    RemoteString string `json:"fs,omitempty"`
 
}



// List the files in a given remote/path
type OperationsList struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
    ListOptions map[string]string `json:"opt,omitempty"`
 
}



// Create a folder on the remote
type OperationsMkdir struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
 
}



// Move a file or directory from source remote to destination remote
type OperationsMovefile struct {
    JsonData string `json:"-"`
    SourceRemote string `json:"srcFs,omitempty"`
    SourcePath string `json:"srcRemote,omitempty"`
    DestinationRemote string `json:"dstFs,omitempty"`
    DestinationPath string `json:"dstRemote,omitempty"`
 
}



// Create a public link to the remote path (if possible)
type OperationsPubliclink struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
 
}



// Remove a directory/container and all of its contents
type OperationsPurge struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
 
}



// Remove an empty dir/container
type OperationsRmdir struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
 
}



// Remove all empty directories in the path
type OperationsRmdirs struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
    LeaveRoot bool `json:"leaveRoot,omitempty"`
 
}



// Count the number of bytes and files in the remote
type OperationsSize struct {
    JsonData string `json:"-"`
    RemotePath string `json:"fs,omitempty"`
 
}



// Get information about the supplied file or directory
type OperationsStat struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
    StatOptions map[string]string `json:"opt,omitempty"`
 
}



// Upload a file using multipart/form-data
type OperationsUploadfile struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    Path string `json:"remote,omitempty"`
    File string `json:"file,omitempty"`
    UploadOptions map[string]string `json:"opt,omitempty"`
 
}
