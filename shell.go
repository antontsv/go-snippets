package main

import (
	"log"
	"os"
)

func shell() {

	cwd, err := os.Getwd()
	checkErr(err)

	attributes := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}
	proc, err := os.StartProcess("/usr/bin/env", []string{"bash"}, attributes)
	checkErr(err)

	_, err = proc.Wait()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
