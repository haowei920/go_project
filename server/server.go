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

	http.HandleFunc("/hello", HandleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	responseData, err := ioutil.ReadAll(r.Body)
	var requestParamObject request_param
	var dataObject response_filenames

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = json.Unmarshal(responseData, &requestParamObject)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result, _ := returnFiles(&(requestParamObject.Name), &(requestParamObject.Path))
	dataObject.Data = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataObject)

}

func returnFiles(regexPtr *string, pathPtr *string) ([]string, error) {
	*regexPtr = "^" + *regexPtr + "$"

	fmt.Println("regex:", *regexPtr)
	fmt.Println("path:", *pathPtr)

	re := regexp.MustCompile(*regexPtr)

	fmt.Println("This is re:", re)
	var result []string
	err := filepath.Walk(*pathPtr,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if re.MatchString(path) {
				// fmt.Println(path)
				result = append(result, path)

			}

			return nil
		})
	if err != nil {
		// log.Println(err)
		fmt.Println("This path cannot be found")
		return nil, err
	}
	return result, nil
}
