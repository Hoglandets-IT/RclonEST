package endpoints 


// Set the bandwidth limit in use
type CoreBwlimit struct {
    JsonData string `json:"-"`
    Rate string `json:"rate,omitempty"`
 
}



// Run a rclone terminal command over rc
type CoreCommand struct {
    JsonData string `json:"-"`
    Command string `json:"command,omitempty"`
    Arguments []string `json:"arg,omitempty"`
    Options map[string]string `json:"opt,omitempty"`
    ReturnType string `json:"returnType,omitempty"`
 
}



// Runs a garbage collection
type CoreGc struct {
    JsonData string `json:"-"`
 
}



// List all the stats groups currently in memory
type CoreGroupList struct {
    JsonData string `json:"-"`
 
}



// Prints the memory stats
type CoreMemstats struct {
    JsonData string `json:"-"`
 
}



// Obscure a passed in string
type CoreObscure struct {
    JsonData string `json:"-"`
    Clear string `json:"clear,omitempty"`
 
}



// Return the PID of the current process
type CorePid struct {
    JsonData string `json:"-"`
 
}



// Shut down the rclone server
type CoreQuit struct {
    JsonData string `json:"-"`
    ExitCode int `json:"exitCode,omitempty"`
 
}



// Prints the stats information
type CoreStats struct {
    JsonData string `json:"-"`
    Group string `json:"group,omitempty"`
 
}



// Deletes the stats group
type CoreStatsDelete struct {
    JsonData string `json:"-"`
    Group string `json:"group,omitempty"`
 
}



// Resets the counters, errors and finished transfers for a given group or all
type CoreStatsReset struct {
    JsonData string `json:"-"`
    Group string `json:"group,omitempty"`
 
}



// Return stats about finished transfers
type CoreTransferred struct {
    JsonData string `json:"-"`
    Group string `json:"group,omitempty"`
 
}



// Shows the current version of rclone and the go runtime
type CoreVersion struct {
    JsonData string `json:"-"`
 
}
