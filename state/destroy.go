package state

import "os"

func Destroy() error {
	return os.Remove("./.dcont/state.json")
}
