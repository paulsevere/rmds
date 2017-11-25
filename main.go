// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/paulsevere/rmds/walker"
)

func main() {
	args := os.Args[1:]
	// spew.Dump(args[0])
	dir, err := os.Getwd()
	if err != nil {
		println(err.Error())
		return
	}
	var target_dir string
	var target_rm string

	if len(args) == 1 {
		target_dir = path.Join(dir, args[0])
		target_rm = ".DS_Store"
	} else if len(args) == 2 {
		target_dir = path.Join(dir, args[1])
		target_rm = args[0]
	} else {
		target_dir = dir
	}
	if pathIsDir(target_dir) {
		walker.Walk(target_rm, target_dir)
	} else {
		color.Red("Please enter a valid directory path")
	}

	// walker.Walk(dir)

}

func pathIsDir(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
