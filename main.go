package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://islandman.org")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// byteSlice is container the resp.Body.Read loads data
	// Assumes that input data is < 100,000 bytes
	// byteSlice := make([]byte, 99999) //
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	// Send Response string to Stdout
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes")
	return len(bs), nil
}
