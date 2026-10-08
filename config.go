package main

type config struct {
	latest version 
	installed versionLocal
	cursorUrl string 
}

type version struct {
	version string
	url string
}

type versionLocal struct {
	version string
	path string
}