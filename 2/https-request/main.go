package main

import "net/http"

func main() {

	req, err := http.Get("http:/google.com")
	if err != nil {
		panic(err)
	}
}
