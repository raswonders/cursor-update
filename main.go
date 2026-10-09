package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	path := flag.String("path", os.Getenv("CURSOR_APPIMAGE"), "path to the installed Cursor AppImage")
	platform := flag.String("platform", "linux-x64", "download platform")
	track := flag.String("track", "stable", "release track")
	flag.Parse()

	if *path == "" {
    fmt.Fprintln(os.Stderr, "path is required, pass it via flag or env. var CURSOR_APPIMAGE")
    flag.Usage()
    os.Exit(2)
  }

	config := config{
		cursorUrl: fmt.Sprintf("https://cursor.com/api/download?platform=%s&releaseTrack=%s", *platform, *track),
		installed: versionLocal{
			path: *path,
		},
		latest: version{},
	}

  needsUpdate, err := isInstalledBehind(&config) 
	if err != nil {
		fmt.Printf("Couldn't verify versions: %v\n", err)
		os.Exit(1)
	}

	if needsUpdate {
		err := fetchLatestAppImage(&config)
		if err != nil {
			fmt.Printf("Couldn't download cursor's image: %v\n", err)
		}
		fmt.Println("Cursor was downloaded.")

		src := "cursor.AppImage"
		dst := config.installed.path

		if _, err := os.Stat(dst); err == nil {
			if err := os.Rename(dst, dst+".bak"); err != nil {
				fmt.Printf("Couldn't create backup for installed image: %v\n", err)
				os.Exit(1)
			}
		}

		if err := os.Rename(src, dst); err != nil {
			fmt.Printf("Couldn't move latest image to destination: %v\n", err)
			os.Exit(1)
		} 

		fmt.Printf("Latest cursor.AppImage (v%s) was installed\n", config.latest.version)
		os.Exit(0)
	}
	fmt.Printf("Latest version (v%s) already installed.\n", config.installed.version)
}