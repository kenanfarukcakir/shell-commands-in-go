package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var hidden bool
var debug bool
var path string

func init() {
	flag.BoolVar(&hidden, "hidden", false, "show hidden files also")
	flag.BoolVar(&debug, "debug", false, "show debug logs about executable, current path...")
	flag.StringVar(&path, "path", "", "path")
}

// ls -- list directory contents
func main() {
	flag.Parse()
	var err error

	if debug {
		pName := os.Args[0]
		fmt.Printf("\nExecutable name is %s \n", pName)

		for i, arg := range os.Args {
			fmt.Printf("Arg %d : %s ", i, arg)
		}

		currPath, err := os.Executable()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("\nExecutable lies in %s \n", currPath)
	}

	if path == "" { // specific path not given
		path, err = os.Getwd() // set current directory
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// check if path given contains any '~'
		if path[0] == '~' {
			builder := strings.Builder{}
			builder.WriteString(os.Getenv("HOME"))
			builder.WriteString(path[1:])
			path = builder.String()
		}

	}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	builder := strings.Builder{}

	hiddenColor := color.New(color.Underline)
	dirColor := color.New(color.FgBlue)
	fileColor := color.New(color.FgWhite)

	for _, entry := range dirEntries {
		if entry.Name()[0] == '.' { // means it is hidden
			if hidden { // show hiddens
				builder.WriteString(hiddenColor.Sprintf("%s ", entry.Name()))
			}
			continue
		}
		if entry.IsDir() {
			builder.WriteString(dirColor.Sprintf("%s ", entry.Name()))
		} else if entry.Type().IsRegular() {
			builder.WriteString(fileColor.Sprintf("%s ", entry.Name()))
		}

		// add executable file check

	}

	fmt.Println(builder.String())

}
