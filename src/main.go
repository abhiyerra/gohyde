package main

import (
	"fmt"
	"io/ioutil"
	"flag"
)

func Encode(path *string) string {
	var file []byte
	var err error

	if file, err = ioutil.ReadFile(path) {
		return len(file)
	}

	return string(file)
}

func Decode(url *string) void {
	
}

func main() {
	var filename = flag.String("f", "", "filename of the guy")
	flag.Parse()
	
	Encode(filename)

        fmt.Println("Hello, world.  Sqrt(2) = ", *filename)
	
}