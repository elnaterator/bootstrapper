package core

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetJson(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		panic(err)
	}
	return result
}
