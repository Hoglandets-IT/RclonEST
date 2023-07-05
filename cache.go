package main

type CacheExpire struct {
	RCPath   bool   `json:"-" rcpath:"/cache/expire"`
	JsonData string `json:"-"`
	Remote   string `json:"remote"`
	WithData bool   `json:"withData"`
}

type CacheFetch struct {
	RCPath   bool   `json:"-" rcpath:"/cache/fetch"`
	JsonData string `json:"-"`
}

type CacheStats struct {
	RCPath   bool   `json:"-" rcpath:"/cache/stats"`
	JsonData string `json:"-"`
}
