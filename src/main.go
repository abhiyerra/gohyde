package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}

func encode(path *string) (content []string, err error) {
	filepath := *path
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	base64File := base64.StdEncoding.EncodeToString(file)
	res := ""
	var fileslice []string
	for i, r := range base64File {
		res = res + string(r)
		if i > 0 && (i+1)%140 == 0 {
			fileslice = append(fileslice, res)
			base64File = removeCharacters(base64File, res)
		}
	}
	fileslice = append(fileslice, base64File)
	return fileslice, nil
}

func decode(url *string) {
}

func main() {
	var filename = flag.String("f", "", "filename of the guy")
	flag.Parse()

	file, _ := encode(filename)
	for _, k := range file {
		fmt.Print(k)
	}
}
