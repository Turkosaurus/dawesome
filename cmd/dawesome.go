package dawesome

import (
	"fmt"
)

type bar struct {
	author string
	count  int
	data   []byte
}

func PlayBar() (err error) {
	var bar bar
	bar.author = "Turk"
	bar.count = 4
	fmt.Println(bar)
	return nil
}
