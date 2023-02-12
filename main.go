package main

import (
	"github.com/debdutdeb/devcontainer-lite/cmd"
)

/*
- Build the image
- Start the container
- set up a timer to commit the image depending on fs diff
- run it in the bg (save state in current directory to dictate what to do)
- start neovim as you would?
*/

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
