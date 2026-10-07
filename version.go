package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func findInstalledVersion() (string, error) {
	prefix := "X-AppImage-Version="
	out, err := exec.Command("7z", "e", "-so", "/opt/cursor/cursor.AppImage", "cursor.desktop").Output()
	if err != nil {
		return "", err
	} 

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		version, ok := strings.CutPrefix(line, prefix)	
		if ok {
			return version, nil
		}
	}

	return "", fmt.Errorf("Couldn't find installed version");
}

func findLatestVersion() (string, error) {
	data, err := fetchDownloadData()	
	if err != nil {
		return "", err
	}

	return data.Version, nil
}

func isInstalledBehind() (bool, error) {
	installed, err := findInstalledVersion()
	if err != nil {
		return false, err
	}
	latest, err := findLatestVersion()
	if err != nil {
		return false, err
	}

	return isVersionLess(installed, latest) 
}

func isVersionLess(a, b string) (bool, error) {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	n := max(len(as), len(bs))
	for i := range n {
		var ai, bi int
		var err error
		if i < len(as[i]) {
			ai, err = strconv.Atoi(as[i])
			if err != nil {
				return false, err
			}
		}

		if i < len(bs[i]) {
			bi, err = strconv.Atoi(bs[i])
			if err != nil {
				return false, err
			}
		}

		if ai != bi {
			return ai < bi, nil
		}
	}
	return false, nil
}