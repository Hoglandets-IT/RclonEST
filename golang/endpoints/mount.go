package endpoints 


// List all mounted remotes
type MountListmounts struct {
    JsonData string `json:"-"`
 
}



// Create a new mount point
type MountMount struct {
    JsonData string `json:"-"`
    Remote string `json:"fs,omitempty"`
    MountPoint string `json:"mountPoint,omitempty"`
    MountType string `json:"mountType,omitempty"`
    MountOptions string `json:"mountOpt,omitempty"`
    VfsOptions string `json:"vfsOpt,omitempty"`
 
}



// Show all possible mount types
type MountTypes struct {
    JsonData string `json:"-"`
 
}



// Unmount a mount point
type MountUnmount struct {
    JsonData string `json:"-"`
    MountPoint string `json:"mountPoint,omitempty"`
 
}



// Unmount all mount points
type MountUnmountall struct {
    JsonData string `json:"-"`
 
}
