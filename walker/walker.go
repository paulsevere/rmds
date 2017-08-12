package walker

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/michaelTJones/walk"
)

var count = 0

func Walk(path string) {
	walk.Walk(path, walker)
	colorPrint(count)
}

func walker(path string, info os.FileInfo, err error) error {
	if info.Name() == ".DS_Store" {
		// spew.Dump(path)
		delerr := os.Remove(path)
		if delerr != nil {
			return delerr

		} else {
			count++

			// println("Removed DS_Store in " + path)
		}

	}
	// spew.Dump(path, info.IsDir())
	return err
}

func colorPrint(i interface{}) {
	yellow := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("Removed %v items", yellow(i))
}
