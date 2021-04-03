package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

	requestBytes, _ := json.Marshal(requestParamObject)

	url := "http://localhost:8080/hello"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	var dataObject response_filenames
	err = json.Unmarshal(body, &dataObject)
	result1 := strings.Join(dataObject.Data, "\n")
	fmt.Println("response Body:", result1)

}
