package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// Custom Datatype with Embedded Mutex Serving as Spoof Database for WebSite Maps
type WebDirectory struct {
	sync.Mutex
	webMap map[string]string
}

// Custom Datatype for Regularized Responses
type regularResponse struct {
	StstusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Error      string      `json:eror`
	Data       interface{} `json:"data"`
}

var webDirectoryObj = WebDirectory{sync.Mutex{}, make(map[string]string)}

func main() {
	go webStatusChecker()

	mux := http.DefaultServeMux
	mux.HandleFunc("POST /websites", addWebsites)
	mux.HandleFunc("GET /websites", webStatus)
	fmt.Println("Server Running on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// Website Checker CRONJOB function
func webStatusChecker() {
	for true {
		for key := range webDirectoryObj.readDirectory() {
			res, Fetcherr := http.Get("http://" + key)
			if Fetcherr != nil {
				fmt.Println("Error Fetching from", key)
				webDirectoryObj.writeDirectory(key, "INVALID")
			} else {
				if res.StatusCode == 200 {
					webDirectoryObj.writeDirectory(key, "UP")
				} else {
					webDirectoryObj.writeDirectory(key, "DOWN")
				}
			}
		}
		time.Sleep(time.Second * 5)
	}
}

// POST /websites Handler
func addWebsites(res http.ResponseWriter, req *http.Request) {
	var response regularResponse
	fmt.Println("Incoming Post Request On /websites")
	body, bodyError := io.ReadAll(req.Body)

	if bodyError != nil {
		response.assembleResult(http.StatusBadGateway, "Error reading request body", bodyError.Error(), nil, res)
		return
	}

	var tempmap = make(map[string][]string)
	json.Unmarshal(body, &tempmap)
	for _, key := range tempmap["websites"] {
		webDirectoryObj.writeDirectory(key, "UNCHECKED")
	}
	response.assembleResult(http.StatusAccepted, "Websites Added Successfully", "", nil, res)
}

// GET /websites Handler
func webStatus(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming Get Request On /websites")
	var response regularResponse
	queryParam := req.URL.Query().Get("param")

	if queryParam != "" {
		// Filter the result based on the query parameter
		filteredResult, err := webDirectoryObj.filterDirectory(queryParam)
		if err != nil {
			response.assembleResult(http.StatusNotFound, "LookUp Error", err.Error(), nil, res)
			return
		}
		response.assembleResult(http.StatusAccepted, "Request Completed Successfully", "", filteredResult, res)
		return
	}
	// Default Map Return
	response.assembleResult(http.StatusAccepted, "Request Completed Successfully", "", webDirectoryObj.readDirectory(), res)
	return
}

// Reads the WebDirectory
func (directoryObj WebDirectory) readDirectory() map[string]string {
	directoryObj.Lock()
	defer directoryObj.Unlock()
	return directoryObj.webMap
}

// Writes the new values WebDirectory WebMap
func (directoryObj WebDirectory) writeDirectory(key, value string) {
	directoryObj.Lock()
	defer directoryObj.Unlock()
	directoryObj.webMap[key] = value
	return
}

// Returnes Value Associated with the Queried key Error Message if Absent
func (directoryObj WebDirectory) filterDirectory(queryParam string) (string, error) {
	directoryObj.Lock()
	defer directoryObj.Unlock()
	result, isPresent := directoryObj.webMap[queryParam]
	if !isPresent {
		return "", fmt.Errorf("Provided key %v is not present in the database, kindly call the POST API prior to querying", queryParam)
	}
	return fmt.Sprintf("%v is %v", queryParam, result), nil
}

// Assembles the Result in a Regularized Format
func (response *regularResponse) assembleResult(statusCode int, message string, err string, data interface{}, res http.ResponseWriter) {
	*&response.StstusCode, *&response.Message, *&response.Error, *&response.Data = statusCode, message, err, data
	out, MarshalErr := json.Marshal(*response)

	if MarshalErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("JSON Marshal Error"))
		return
	}
	res.WriteHeader(statusCode)
	res.Write(out)
	return
}
