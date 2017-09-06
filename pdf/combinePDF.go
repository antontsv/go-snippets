package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

}

func checkErr(msg string, err error) {
	if err != nil {
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprintf(os.Stderr, msg)
		os.Exit(1)
	}
}
