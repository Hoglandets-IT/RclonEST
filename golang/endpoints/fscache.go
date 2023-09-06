package endpoints 


// Clears the filesystem cache
type FscacheClear struct {
    JsonData string `json:"-"`
 
}



// Returns the number of entries in the fs cache
type FscacheEntries struct {
    JsonData string `json:"-"`
 
}
