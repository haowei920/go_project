package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//request_param struct that contains regex name of file (Name) to be searched and directory to search recrusively (Path)
type request_param struct {
	Name string
	Path string
}

//response_filenames contains string array with all filenames that fullfills the regex file name to be searched in the directory given
type response_filenames struct {
	Data []string
}

func main() {
	regexPtr := flag.String("name", "ECE411", "Regular expression for the file name you are trying to locate")
	pathPtr := flag.String("path", ".pdf", "Directory in which you are searching for the file name")

	flag.Parse()
	//instantiating and populating request_param object using inputs from user
	var requestParamObject request_param
	requestParamObject.Name = *regexPtr
	requestParamObject.Path = *pathPtr

	//encode requestParamObject to a JSON string
	requestBytes, err := json.Marshal(requestParamObject)
	if err != nil {
		log.Fatal(err)
	}

	url := "http://localhost:8080/find-file"

	//for control over HTTP client headers, create a Client
	client := &http.Client{}
	// making http post request to server with our json data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//client must close the response body when finished with it
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {

		var dataObject response_filenames
		// to convert json data back into our array of struct before printing it out
		err = json.Unmarshal(body, &dataObject)
		result1 := strings.Join(dataObject.Data, "\n")
		fmt.Println(result1)
	} else {
		log.Fatalf("Bad status code %d received: %s", resp.StatusCode, string(body))

	}

}
