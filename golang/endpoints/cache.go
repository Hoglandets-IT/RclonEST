package endpoints 


// Purge a remote from the cache backend. Supports either directory or file
type CacheExpire struct {
    JsonData string `json:"-"`
    Remote string `json:"remote,omitempty"`
    WithData bool `json:"withData,omitempty"`
 
}



// Ensure the specified cchinks are cached on disk
type CacheFetch struct {
    JsonData string `json:"-"`
 
}



// Print cache statistics
type CacheStats struct {
    JsonData string `json:"-"`
 
}
