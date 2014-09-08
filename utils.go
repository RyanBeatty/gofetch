package gofetch

import (
	"fmt"
	"os"
)

func RaiseOnError(err error, message string) {
	if err != nil {
		fmt.Printf("%s\n", message)
		os.Exit(1)
	}
}