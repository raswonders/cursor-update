package main

import (
	"fmt"
	"time"
)

type Spinner struct {
	Frames []string
	FPS    time.Duration
}

var DotSpinner = Spinner{
	Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
	FPS:    time.Second / 10,
}

const (
	hideCursor = "\033[?25l"
	showCursor = "\033[?25h"
	cleanLine  = "\033[2K"
)

func ShowSpinner(spinner Spinner, msg string, stop <-chan struct{}) {
	ticker := time.NewTicker(spinner.FPS)
	defer ticker.Stop()
	defer fmt.Printf(showCursor)
	defer fmt.Printf("\r%s", cleanLine)

	i := 0
	fmt.Println()
	fmt.Printf(hideCursor)
	for {
		select {
		case <-stop:
				return
		case <-ticker.C:
			fmt.Printf("\r%s%s", spinner.Frames[i], msg)
			i = (i + 1) % len(spinner.Frames)
		}
	}
}	
