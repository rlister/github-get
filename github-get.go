package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

// error handler
func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

// create a closure to get from repo with our creds
func client(repo string, token string) func(string) []byte {
	client := &http.Client{}

	// return a get() function
	return func(src string) []byte {
		url := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, src)

		request, err := http.NewRequest("GET", url, nil)
		check(err)

		request.Header.Add("Accept", "application/vnd.github.v3.raw")
		if len(token) > 0 {
			request.Header.Add("Authorization", fmt.Sprintf("token %s", token))
		}
		response, err := client.Do(request)
		check(err)

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		check(err)

		return body
	}
}

func main() {
	get := client(os.Getenv("REPO"), os.Getenv("TOKEN"))

	// loop arguments
	for _, arg := range os.Args[1:] {

		// optionally request output file with src:dest
		files := strings.SplitN(arg, ":", 2)
		output := get(files[0])

		// write to requested file, or stdout
		if len(files) > 1 {
			// mkpath
			err := os.MkdirAll(path.Dir(files[1]), 0755)
			// parse json
			var dir []interface{}
			err := json.Unmarshal(dirname, &dir)
			check(err)

			// write the file
			err = ioutil.WriteFile(files[1], output, 0644)
			check(err)
		} else {
			fmt.Println(string(output))
		}
	}
}
