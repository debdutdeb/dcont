package actions

import (
	"errors"
	"fmt"
	"go/build"
	"os"
)

type _languageServerOpt struct {
	startUpCommand string
	mounts         []string
}

type languageServerOpt map[string]_languageServerOpt

var languageServerOpts languageServerOpt = make(languageServerOpt)

func init() {
	var (
		goPath  string
		currDir string
	)
	goPath, ok := os.LookupEnv("GOPATH")
	if !ok {
		goPath = build.Default.GOPATH
	}
	currDir, _ = os.Getwd()
	mounts := []string{
		fmt.Sprintf("%s:%s", goPath, goPath),
		fmt.Sprintf("%s:%s", currDir, currDir),
	}
	languageServerOpts["golang"] = _languageServerOpt{
		startUpCommand: "gopls -vv --port %d &>/gopls.log &",
		mounts:         mounts,
	}
}

func (l languageServerOpt) Get(lang string) _languageServerOpt {
	return l[lang]
}

func IsGolang() bool {
	if _, err := os.Stat("go.mod"); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
