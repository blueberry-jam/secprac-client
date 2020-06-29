package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Script represents a vulnerability-checking script provided by the server
type Script struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Shell       string `json:"shell"`
	Script      string `json:"script"`
	URL         string `json:"url"`
}

// GetScripts fetches the vulnerability-checking scripts from the specified remote server\
func GetScripts(remote, token string) ([]Script, error) {

	// Send GET request
	client := &http.Client{}
	req, err := http.NewRequest("GET", remote+"/vulns/vuln.json", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("server responded with bad status code: " + strconv.Itoa(resp.StatusCode))
	}

	// Read response data
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON into slice
	var scripts []Script
	err = json.Unmarshal(body, &scripts)
	if err != nil {
		return nil, err
	}

	return scripts, nil
}