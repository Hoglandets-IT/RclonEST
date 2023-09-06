package endpoints 


// Set runtime.SetBlockProfileRate for blocking profiling
type DebugSetBlockProfileRate struct {
    JsonData string `json:"-"`
    Rate int `json:"rate,omitempty"`
 
}



// Call runtime/debug.SetGCPercent for setting the garbage collection target percentage.
type DebugSetGcPercent struct {
    JsonData string `json:"-"`
    Rate int `json:"rate,omitempty"`
 
}



// Set runtime.SetMutexProfileFraction for mutex profiling
type DebugSetMutexProfileFraction struct {
    JsonData string `json:"-"`
    Rate int `json:"rate,omitempty"`
 
}



// Call runtime/debug.SetMemoryLimit for setting a soft memory limit for the runtime.
type DebugSetSoftMemoryLimit struct {
    JsonData string `json:"-"`
    MemLimit string `json:"mem-limit,omitempty"`
 
}
