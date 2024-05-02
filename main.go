package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
  // Creates a new io.Reader 
	r := strings.NewReader("Hello, Reader!")

  // create a buffer to store
	b := make([]byte, 8)

  // read
	for {
		n, err := r.Read(b)
    fmt.Printf("val:\"%v\" buffer:%v bytes:%v error:%v \n",string(b[:n]),b[:n],n,err)
		if err == io.EOF {
			break
		}
	}
}
