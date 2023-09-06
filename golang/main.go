package main

import "fmt"

func main() {
	clint := NewClient("https://postman-echo.com", "admin", "admin")

	resp, err := clint.postRequest("/post", "{\"json_data\": {\"goodbye\": \"world\"}, \"hello\":\"world\"}")
	fmt.Println(resp)
	fmt.Println(err)

}
