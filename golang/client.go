package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

type Client interface {
	makeRequest(req *http.Request) ([]byte, error)
	postRequest(path string, body string) ([]byte, error)
	Run(command interface{}) ([]byte, error)
}

type RClient struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient *http.Client
}

type Operation struct {
	JsonData string `json:"-"`
}

func (meta Operation) GetJsonData() string {
	return meta.JsonData
}

type RCCredentials struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewClient(baseURL, username, password string) *RClient {
	var client *RClient
	client = &RClient{
		BaseURL:    baseURL,
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}
	return client
}

func ClientFromJson(json_data string) *RClient {
	var clientJ *RCCredentials
	Unmarshal(json_data, &clientJ)
	return NewClient(clientJ.Url, clientJ.Username, clientJ.Password)
}

func getRCPathTag(command *interface{}) (string, error) {
	// Check if type is interface or pointer
	tpeofx := reflect.TypeOf(command)
	if tpeofx.Kind() == reflect.Ptr {
		vofx := reflect.ValueOf(command)
		if vofx.IsNil() {
			return "", fmt.Errorf("Nil pointer")
		}
		tpeofx := tpeofx.Elem()
		if tpeofx.Kind() != reflect.Struct {
			return "", fmt.Errorf("Not a struct")
		}
		valofx := reflect.ValueOf(command).Elem().FieldByName("RCPath")

		if !valofx.IsValid() {
			return "", fmt.Errorf("No RCPath field")
		}
	} else {
		_, ok := tpeofx.FieldByName("RCPath")
		if !ok {
			return "", fmt.Errorf("No RCPath field")
		}
	}

	field, ok := tpeofx.FieldByName("RCPath")
	if !ok {
		return "", fmt.Errorf("No RCPath field")
	}

	rcpath := field.Tag.Get("rcpath")
	if rcpath == "" {
		return "", fmt.Errorf("No rcpath tag")
	}
	return rcpath, nil
}

func funcToUrl(command *interface{}) string {
	rcpath, err := getRCPathTag(command)
	fmt.Println("Command: ", command)
	fmt.Println("RCP: ", rcpath)
	x := reflect.TypeOf(command)
	fmt.Println(x)
	if err != nil {
		funcName := reflect.TypeOf(command).Name()
		if funcName == "" {
			funcName = reflect.TypeOf(command).Elem().Name()
		}
		snake := matchFirstCap.ReplaceAllString(funcName, "${1}/${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}/${2}")
		return strings.ToLower(snake)
	}
	return rcpath
}

func joinUrl(url, path string) string {
	if strings.HasSuffix(url, "/") {
		return url + path
	}
	return url + "/" + path
}

func (s *RClient) makeRequest(req *http.Request) ([]byte, error) {
	// req.SetBasicAuth(s.Username, s.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("Unauthorized")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *RClient) postRequest(path string, body string) ([]byte, error) {
	bodyReader := strings.NewReader(body)

	req, err := http.NewRequest("POST", joinUrl(s.BaseURL, path), bodyReader)
	if err != nil {
		return nil, err
	}

	return s.makeRequest(req)
}

func (s *RClient) Run(command *interface{}) ([]byte, error) {
	var jsonCommand string = ""
	// cType := reflect.ValueOf(command).Elem().Type
	// a := cType.FieldByName("JsonData")
	// if a.IsValid() {
	// 	jsonCommand = a.String()
	// }

	if jsonCommand == "" {
		marshJsonCommand, err := json.Marshal(command)
		if err != nil {
			return nil, err
		}
		jsonCommand = string(marshJsonCommand)
	}

	rcPathTag := funcToUrl(command)
	fmt.Println(command, rcPathTag, jsonCommand)
	return s.postRequest(rcPathTag, jsonCommand)
}

func (s *RClient) RunC(path string, commandLn string) ([]byte, error) {
	fmt.Println("RunC: ", path, commandLn)
	return s.postRequest(path, commandLn)
}
