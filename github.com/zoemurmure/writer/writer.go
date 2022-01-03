package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type customWriter struct{}

type GithubRespond []struct {
	FullName string `json:"full_name"`
}

func (cw customWriter) Write(p []byte) (n int, err error) {
	var resp GithubRespond
	json.Unmarshal(p, &resp)
	for _, v := range resp {
		fmt.Println(v.FullName)
	}
	return len(p), nil
}

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		panic(err.Error())
	}

	var writer customWriter
	io.Copy(writer, resp.Body)
}
