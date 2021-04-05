package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type request_param struct {
	Name string
	Path string
}

type response_filenames struct {
	Data []string
}

func main() {
	//tells http packaeg to handle all requests to /hello with HandleRequest
	http.HandleFunc("/hello", HandleRequest)
	//listen to port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// http.ResponseWriter is a mechanism used for sending responses to any connected HTTP clients (return to client)
// http.Request is how we retrieve data from web request (from client)
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	//responseData contains our data from client
	responseData, err := ioutil.ReadAll(r.Body)
	var requestParamObject request_param
	var dataObject response_filenames

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// populate our request_param object using data from client
	err = json.Unmarshal(responseData, &requestParamObject)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//call returnFiles function that takes regex of file name and direcotry to search for and return array of file names that fulfills the criteria
	result, _ := returnFiles(&(requestParamObject.Name), &(requestParamObject.Path))
	dataObject.Data = result

	w.WriteHeader(http.StatusOK)
	//populates our return http.ResponseWriter with dataObject and in json format
	json.NewEncoder(w).Encode(dataObject)

}

func returnFiles(regexPtr *string, pathPtr *string) ([]string, error) {
	//ensure that there is nothing before and after our regex expression
	*regexPtr = "^" + *pathPtr + *regexPtr + "$"
	//ensures expression can be parsed
	re := regexp.MustCompile(*regexPtr)

	var result []string
	err := filepath.Walk(*pathPtr,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//add to result if matches with all file path found recursively in given directory
			if re.MatchString(path) {
				result = append(result, path)

			}

			return nil
		})
	if err != nil {
		fmt.Println("This path cannot be found")
		return nil, err
	}
	return result, nil
}
