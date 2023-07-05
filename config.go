package main

type ConfigCreate struct {
	JsonData   string            `json:"-"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Parameters map[string]string `json:"parameters"`
	Options    map[string]string `json:"options"`
}

type ConfigDelete struct {
	JsonData string `json:"-"`
	Name     string `json:"name"`
}

type ConfigDump struct {
	JsonData string `json:"-"`
}

type ConfigGet struct {
	Operation `json:"-"`
	Name      string `json:"name"`
}

type ConfigListremotes struct {
	JsonData string `json:"-"`
}

type ConfigProviders struct {
	JsonData string `json:"-"`
}

type ConfigUpdate struct {
	JsonData   string            `json:"-"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
	Options    map[string]string `json:"options"`
}
