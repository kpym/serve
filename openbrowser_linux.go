package main

import (
	"os/exec"
)

// open the browser
func openbrowser(url string) error {
	return exec.Command("xdg-open", url).Start()
}
