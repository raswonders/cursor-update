package main

import (
	"fmt"
	"os/exec"
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