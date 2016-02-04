package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wojciechp/projectutils/gitinit"
)

var path = flag.String("local", "", "Path to local git repository (will be created if necessary)")
var remote = flag.String("add-origin", "", "URL to remote repo to add as git remote origin")

func main() {

	flag.Parse()

	if *path == "" {
		fmt.Printf("Provide a path to the place where local repo should be created\n")
		os.Exit(1)
	}
	if err := os.MkdirAll(*path, os.ModePerm); err != nil {
		fmt.Printf("Directory creation failed: %s", err)
		os.Exit(1)
	}
	git := gitinit.LocalGit{Path: *path}
	fmt.Print("Initializing repo...\n")
	if err := git.Init(); err != nil {
		fmt.Printf("Failure during repo initialization: %s", err)
	}

	if *remote == "" {
		fmt.Printf("Origin not specified - not adding")
	} else {
		if err := git.AddRemote("origin", *remote); err != nil {
			fmt.Printf("Failed to add remote: %s", err)
		}
	}
}
