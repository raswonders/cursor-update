package main

import (
	"fmt"
	"os"
)

func main() {
	config := config{
		cursorUrl: "https://cursor.com/api/download?platform=linux-x64&releaseTrack=stable",
		installed: versionLocal{
			path: "/home/rhepner/Applications/cursor/cursor.AppImage",
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

		src := "cursor.AppImage"
		dst := config.installed.path

		if _, err := os.Stat(dst); err == nil {
			if err := os.Rename(dst, dst+".bak"); err != nil {
				fmt.Printf("Couldn't create backup for installed image: %v", err)
				return
			}
		}

		if err := os.Rename(src, dst); err != nil {
			fmt.Printf("Couldn't move latest image to destination: %v", err)
			return
		} 

		fmt.Printf("Latest cursor.AppImage (v%s) was installed", config.latest.version)
	}
}