package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintSomething2(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go PrintSomething2("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	consoleOuput := string(result)

	os.Stdout = stdOut
	if !strings.Contains(consoleOuput, "epsilon") {
		t.Errorf("Expected to find epsilon but it is not there")
	}
}
