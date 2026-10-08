package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type data struct {
	DownloadUrl string	`json:"downloadUrl"` 
	Version string			`json:"version"`
} 

func fetchLatestData(conf *config) error {
	res, err := http.Get(conf.cursorUrl)
	if err != nil {
		return err
	}

	dataRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data data;
	json.Unmarshal(dataRaw, &data)
	conf.latest = version{
		version: data.Version,
		url: data.DownloadUrl,
	}
	return nil 
}

func fetchLatestAppImage(conf *config) error {
	res, err := http.Get(conf.latest.url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile("cursor.AppImage", body, 0755)
	if err != nil {
		return err
	}

	return nil
}