package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = checkPrefix(url)
		//fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//content, err := ioutil.ReadAll(resp.Body)
		bytes, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", content)
		fmt.Printf("\nStatus: %v, Has read %v bytes\n", resp.Status, bytes)
	}
}

func checkPrefix(url string) string {
	if strings.HasPrefix(url, "http") {
		return url
	} else {
		return "http://" + url
	}
}
