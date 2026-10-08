package main

import (
	"fmt"
)

func main() {
	config := config{
		cursorUrl: "https://cursor.com/api/download?platform=linux-x64&releaseTrack=stable",
		installed: versionLocal{
			path: "/opt/cursor/cursor.AppImage",
		},
		latest: version{},
	}

  needsUpdate, err := isInstalledBehind(&config) 
	if err != nil {
		fmt.Printf("Couldn't verify versions: %v", err)
		return
	}

	if needsUpdate {
		err := fetchLatestAppImage(&config)
		if err != nil {
			fmt.Printf("Couldn't download cursor's image: %v", err)
		}
		fmt.Println("Cursor was downloaded.")

	}
}