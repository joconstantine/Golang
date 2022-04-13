package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ResponseWriter struct{}

func main() {
	resp, err := http.Get(("http://google.com"))

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	respWriter := ResponseWriter{}
	io.Copy(respWriter, resp.Body)
	io.Copy(os.Stdout, resp.Body)
}

func (rw ResponseWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}