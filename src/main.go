package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"strings"

	"log"

	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
)

func init() {
	consumerKey = os.Getenv("consumerKey")
	consumerSecret = os.Getenv("consumerSecret")
	accessToken = os.Getenv("accessToken")
	accessSecret = os.Getenv("accessSecret")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Environment variables not set!")
	}
}
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

func uploadFile(file string) (err error) {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	_, _, err = client.Statuses.Update(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var filename = flag.String("f", "", "filename of the guy")
	flag.Parse()

	file, _ := encode(filename)
	for _, k := range file {
		if err := uploadFile(k); err != nil {
			log.Printf("Failed to upload file: %s", err)
			break
		}
	}
}
