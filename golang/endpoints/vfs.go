package endpoints 


// Forget files or directories in the directory cache
type VfsForget struct {
    JsonData string `json:"-"`
 
}



// List all the cached files/virtual directories
type VfsList struct {
    JsonData string `json:"-"`
 
}



// Interval to poll the cache for stale objects
type VfsPollInterval struct {
    JsonData string `json:"-"`
    Interval string `json:"interval,omitempty"`
 
}



// Reads the directories for the specified paths and freshens the directory cache
type VfsRefresh struct {
    JsonData string `json:"-"`
 
}



// Prints cache statistics
type VfsStats struct {
    JsonData string `json:"-"`
 
}
