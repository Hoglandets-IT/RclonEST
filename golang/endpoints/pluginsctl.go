package endpoints 


// Add a plugin from an URL
type PluginsctlAddplugin struct {
    JsonData string `json:"-"`
    PluginUrl string `json:"url,omitempty"`
 
}



// Shows all possible plugins by mime type
type PluginsctlGetpluginsfortype struct {
    JsonData string `json:"-"`
    MimeType string `json:"type,omitempty"`
    PluginType string `json:"pluginType,omitempty"`
 
}



// List the loaded plugins
type PluginsctlListplugins struct {
    JsonData string `json:"-"`
 
}



// List the test plugins
type PluginsctlListtestplugins struct {
    JsonData string `json:"-"`
 
}



// Remove a plugin
type PluginsctlRemoveplugin struct {
    JsonData string `json:"-"`
    PluginName string `json:"name,omitempty"`
 
}



// Remove a test plugin
type PluginsctlRemovetestplugin struct {
    JsonData string `json:"-"`
    PluginName string `json:"name,omitempty"`
 
}
