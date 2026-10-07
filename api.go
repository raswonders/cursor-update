package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type data struct {
	DownloadUrl string	`json:"downloadUrl"` 
	Version string			`json:"version"`
} 

func fetchDownloadData() (data, error) {
	url := "https://cursor.com/api/download?platform=linux-x64&releaseTrack=stable"
	res, err := http.Get(url)
	if err != nil {
		return data{}, err
	}

	dataRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return data{}, err
	}

	var data data;
	json.Unmarshal(dataRaw, &data)

	return data, nil 
}