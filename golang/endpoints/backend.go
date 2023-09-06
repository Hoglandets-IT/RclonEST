package endpoints 


// Runs a backend command
type BackendCommand struct {
    JsonData string `json:"-"`
    Command string `json:"command,omitempty"`
    Remote string `json:"fs,omitempty"`
    Arguments []string `json:"arg,omitempty"`
    Options map[string]string `json:"opt,omitempty"`
 
}
