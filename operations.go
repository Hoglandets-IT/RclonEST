package main

type ListOpts struct {
	Recurse       bool     `json:"recurse"`
	NoModTime     bool     `json:"noModTime"`
	ShowEncrpyted bool     `json:"showEncrypted"`
	ShowOrigIDs   bool     `json:"showOrigIDs"`
	ShowHash      bool     `json:"showHash"`
	NoMimeType    bool     `json:"noMimeType"`
	DirsOnly      bool     `json:"dirsOnly"`
	FilesOnly     bool     `json:"filesOnly"`
	Metadata      bool     `json:"metadata"`
	HashTypes     []string `json:"hashTypes"`
}

type OperationsList struct {
	JsonData string   `json:"-"`
	Remote   string   `json:"fs"`
	Path     string   `json:"remote"`
	Opts     ListOpts `json:"opt"`
}

type OperationsMkdir struct {
	JsonData string `json:"-"`
	Remote   string `json:"fs"`
	Path     string `json:"remote"`
}

type OperationsMovefile struct {
	JsonData   string `json:"-"`
	FromRemote string `json:"fromFs"`
	FromPath   string `json:"fromRemote"`
	ToRemote   string `json:"toFs"`
	ToPath     string `json:"toRemote"`
}

type OperationsPubliclink struct {
	JsonData string `json:"-"`
}

type OperationsPurge struct {
	JsonData string `json:"-"`
	Remote   string `json:"fs"`
	Path     string `json:"remote"`
}

type OperationsRmdir struct {
	JsonData string `json:"-"`
	Remote   string `json:"fs"`
	Path     string `json:"remote"`
}

type OperationsRmdirs struct {
	JsonData  string `json:"-"`
	Remote    string `json:"fs"`
	Path      string `json:"remote"`
	LeaveRoot bool   `json:"leaveRoot"`
}

type OperationsSize struct {
	JsonData string `json:"-"`
	FullPath string `json:"fs"`
}

type OperationsStat struct {
	JsonData string `json:"-"`
	Remote   string `json:"fs"`
	Path     string `json:"remote"`
}

type OperationsUploadfile struct {
	JsonData string `json:"-"`
	Remote   string `json:"fs"`
	Path     string `json:"remote"`
}
