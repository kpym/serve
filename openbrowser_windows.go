package main

import (
	"os/exec"
)

// open the browser
func openbrowser(url string) {
	check(exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start())
}
