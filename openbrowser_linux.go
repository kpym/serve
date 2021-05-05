package main

import (
	"os/exec"
)

// open the browser
func openbrowser(url string) {
	check(exec.Command("xdg-open", url).Start())
}
