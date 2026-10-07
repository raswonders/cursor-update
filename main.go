package main

import (
	"fmt"
)

func main() {
	installed, _ := findInstalledVersion()
	latest, _ := findLatestVersion()

	fmt.Println(installed, latest)
}