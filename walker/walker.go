package walker

import (
	"os"

	"github.com/michaelTJones/walk"
)

var count = 0

func Walk(path string) {
	walk.Walk(path, walker)
	println(count)
}

func walker(path string, info os.FileInfo, err error) error {
	count++
	if info.Name() == ".DS_Store" {
		// spew.Dump(path)
		delerr := os.Remove(path)
		if delerr != nil {
			return delerr

		} else {
			println("Removed DS_Store")
		}

	}
	// spew.Dump(path, info.IsDir())
	return err
}
