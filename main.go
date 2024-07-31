package main

import (
	"bytes"
	"encoding/base64"
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
	orgUnitId := args[1]
	orgUnitCode := args[2]
	username := args[3]
	password := args[4]

	makeRequest(url, orgUnitId, orgUnitCode, username, password)
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
