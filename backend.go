package main

type BackendCommand struct {
	JsonData  string            `json:"-"`
	Command   string            `json:"command"`
	Remote    string            `json:"fs"`
	Arguments []string          `json:"arg"`
	Options   map[string]string `json:"opt"`
}
