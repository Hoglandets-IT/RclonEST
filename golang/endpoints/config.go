package endpoints 


// Create a configuration for the remote
type ConfigCreate struct {
    JsonData string `json:"-"`
    Name string `json:"name,omitempty"`
    Parameters map[string]string `json:"parameters,omitempty"`
    Type string `json:"type,omitempty"`
    Options map[string]bool `json:"opt,omitempty"`
 
}



// Delete an existing remote
type ConfigDelete struct {
    JsonData string `json:"-"`
    Name string `json:"name,omitempty"`
 
}



// Dump the configuration to a json object
type ConfigDump struct {
    JsonData string `json:"-"`
 
}



// Get the value for a config entry
type ConfigGet struct {
    JsonData string `json:"-"`
    Name string `json:"name,omitempty"`
 
}



// List all configured remotes
type ConfigListremotes struct {
    JsonData string `json:"-"`
 
}



// Set a password for a remote (deprecated)
type ConfigPassword struct {
    JsonData string `json:"-"`
    Name string `json:"name,omitempty"`
    Parameters map[string]string `json:"parameters,omitempty"`
 
}



// List in JSON format all the providers and options
type ConfigProviders struct {
    JsonData string `json:"-"`
 
}



// Set a path to the config file
type ConfigSetpath struct {
    JsonData string `json:"-"`
    Path string `json:"path,omitempty"`
 
}



// Update the configuration for a remote
type ConfigUpdate struct {
    JsonData string `json:"-"`
    Name string `json:"name,omitempty"`
    Parameters map[string]string `json:"parameters,omitempty"`
    Options map[string]bool `json:"opt,omitempty"`
 
}
