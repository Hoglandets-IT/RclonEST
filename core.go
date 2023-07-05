package main

type CoreBwlimit struct {
	JsonData string `json:"-"`
	Rate     string `json:"rate"`
}

type CoreCommand struct {
	JsonData   string            `json:"-"`
	Command    string            `json:"command"`
	Arguments  []string          `json:"arg"`
	Options    map[string]string `json:"opt"`
	ReturnType string            `json:"returnType"`
}

type CoreGc struct {
	JsonData string `json:"-"`
}

type CoreGroupList struct {
	RCPath   string `rcpath:"/config/group-list"`
	JsonData string `json:"-"`
}

type CoreMemstats struct {
	JsonData string `json:"-"`
}
