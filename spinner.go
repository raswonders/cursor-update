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

func ShowSpinner(spinner Spinner, msg string) {
	fmt.Printf("\033[?25l")
	for i := 0;;i = (i + 1) % len(spinner.Frames) {
		fmt.Printf("\r%s%s", spinner.Frames[i], msg)
		time.Sleep(time.Second / 10)
	}
}
