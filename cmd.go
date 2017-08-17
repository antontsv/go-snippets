package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// filteringWriter is a decorator on top of
// exising writer, but it removes specified search string
// if found before passing it down to a real writer
type filteringWriter struct {
	removeStr string
	w         io.Writer
}

func (f *filteringWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, f.removeStr) {
		s = strings.Replace(s, f.removeStr, "", -1)
	}
	out := []byte(s)
	n, err = f.w.Write([]byte(s))
	return n + len(p) - len(out), err
}

func main() {
	cmd := exec.Command("mycmd", os.Args...)
	cmd.Stdin = strings.NewReader("test")

	outWriter := &filteringWriter{
		removeStr: "Secret",
		w:         os.Stdout,
	}
	errWriter := &filteringWriter{
		removeStr: "Password:",
		w:         os.Stdout,
	}
	cmd.Stdout = outWriter
	cmd.Stderr = errWriter
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
