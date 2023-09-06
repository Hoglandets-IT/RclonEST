package endpoints 


// List all the option blocks
type OptionsBlocks struct {
    JsonData string `json:"-"`
 
}



// Get all the global options
type OptionsGet struct {
    JsonData string `json:"-"`
 
}



// Get currently active config for this call
type OptionsLocal struct {
    JsonData string `json:"-"`
 
}



// Set an option
type OptionsSet struct {
    JsonData string `json:"-"`
 
}
