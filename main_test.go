package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	out := make(chan string)
	go func() {
		var buf strings.Builder
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	f()

	w.Close()
	os.Stdout = originalStdout
	return <-out
}

func TestRootCommand(t *testing.T) {
	output := captureOutput(func() {
		rootCmd.SetArgs([]string{}) // No arguments for root command
		rootCmd.Execute()
	})

	expected := "Hello, World!\n"
	if strings.TrimSpace(output) != strings.TrimSpace(expected) {
		t.Errorf("expected %q, got %q", expected, output)
	}
}
