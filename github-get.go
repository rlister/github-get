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

	// function to get content from github
	get := client(os.Getenv("REPO"), os.Getenv("TOKEN"))

	// loop arguments
	for _, arg := range os.Args[1:] {

		// map of filename to content
		output := make(map[string][]byte)

		// optionally request output file with src:dest
		files := strings.SplitN(arg, ":", 2)
		src := files[0]

		// handle src that is a dir
		if strings.HasSuffix(src, "/") {
			dirname := get(strings.TrimSuffix(src, "/"))

			// parse json
			var dir []interface{}
			err := json.Unmarshal(dirname, &dir)
			check(err)

			// add each file to the output map
			for _, item := range dir {
				file := item.(map[string]interface{})
				name := file["name"].(string)
				if file["type"].(string) == "file" {
					output[name] = get(path.Join(src, name))
				}
			}

		} else { // get src that is a single named file
			output[path.Base(src)] = get(src)
		}

		if len(files) == 1 { // no destination given

			// write each file to stdout
			for _, content := range output {
				fmt.Println(string(content))
			}

		} else { // write to given destination directory

			dest := files[1]

			// make output dir
			err := os.MkdirAll(dest, 0755)
			check(err)

			// write each file to output dir
			for file, content := range output {
				err := ioutil.WriteFile(path.Join(dest, file), content, 0644)
				check(err)
			}

		}
	}
}
