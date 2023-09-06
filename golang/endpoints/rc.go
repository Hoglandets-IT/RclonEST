package endpoints 


// Returns an error with the input as a part of the error string
type RcError struct {
    JsonData string `json:"-"`
 
}



// List all registered remote control commands as a json map
type RcList struct {
    JsonData string `json:"-"`
 
}



// Echo back the input arguments to the output
type RcNoop struct {
    JsonData string `json:"-"`
 
}



// Echo back the input arguments to the output (Requires authentication)
type RcNoopauth struct {
    JsonData string `json:"-"`
 
}
