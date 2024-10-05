package main

import (
	"fmt"
	"runtime"

	"github.com/bchadwic/cheetle/audit"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var impl audit.Auditable
	switch os := runtime.GOOS; os {
	case "windows":
		impl = &audit.Windows{}
	default:
		return fmt.Errorf("os not supported: %s", os)
	}

	impl.LastAccessTime(nil)
	return nil
}
