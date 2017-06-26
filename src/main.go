package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
)

func encode(path *string) (content string, err error) {
	filepath := *path
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	base64File := base64.StdEncoding.EncodeToString(file)
	return base64File, nil
}

func decode(url *string) {
}

func main() {
	var filename = flag.String("f", "", "filename of the guy")
	flag.Parse()

	file, _ := encode(filename)
	fmt.Print(file)
}
