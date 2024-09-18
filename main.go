package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	url := args[0]
	username := args[2]
	password := args[3]
	org_file := args[1]
	records := readCSV(org_file)
	makeRequests(records, url, username, password)
}

func makeRequest(instanceUrl string, orgUnitId string, orgUnitCode string, username string, password string) {
	url := instanceUrl + "/api/organisationUnits/" + orgUnitId
	fmt.Println(url)
	payload := map[string]string{"code": orgUnitCode}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(request.Body)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println(time.Now(), " : ", resp.Status)
	bodyBytes, err := io.ReadAll(resp.Body) // Read the
	if err != nil {
		fmt.Println(err)
	}
	bodyStrings := string(bodyBytes)
	fmt.Println(bodyStrings)
}

func generateCode() string {
	return ""
}

func readCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func makeRequests(records [][]string, instanceUrl string, username string, password string) {
	for _, record := range records {
		orgUnitId := record[1]
		orgUnitCode := record[3]
		makeRequest(instanceUrl, orgUnitId, orgUnitCode, username, password)
	}
}
