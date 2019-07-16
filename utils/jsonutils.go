package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ParseJSON parses json string into a map
func ParseJSON(input io.Reader) (map[string]interface{}, bool) {
	parsed := make(map[string]interface{})
	content, err := ioutil.ReadAll(input)
	if err != nil || json.Unmarshal(content, &parsed) != nil {
		return parsed, false
	}
	return parsed, true
}

// PostJSON wraps http.Post() to reduce boilerplate code.
// Only use this when POST response is not needed
func PostJSON(url string, data interface{}) bool {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}
	return err == nil
}
