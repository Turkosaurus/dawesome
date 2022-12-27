package main

import (
	"fmt"

	dawesome "intonate.dev/dawesome/cmd"
)

func main() {
	err := dawesome.PlayBar()
	if err != nil {
		fmt.Errorf("unknown problem: %w", err)
	}
}
