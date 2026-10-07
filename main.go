package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := fetchDownloadData()	
	if err != nil {
		log.Fatalf("There was an error during fetch: %s", err)
	}

	fmt.Println(data)
}