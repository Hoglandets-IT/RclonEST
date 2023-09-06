package endpoints 


// List all running jobs
type JobList struct {
    JsonData string `json:"-"`
 
}



// Return the status of a job
type JobStatus struct {
    JsonData string `json:"-"`
    JobID int `json:"jobid,omitempty"`
 
}



// Stop a running job
type JobStop struct {
    JsonData string `json:"-"`
    JobID int `json:"jobid,omitempty"`
 
}



// Stop all running jobs in a group
type JobStopgroup struct {
    JsonData string `json:"-"`
    Group string `json:"group,omitempty"`
 
}
