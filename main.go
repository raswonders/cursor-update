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
	fmt.Println(isInstalledBehind(&config))
}