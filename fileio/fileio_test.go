package fileio

import (
    "testing"
	"os"
)


func TestWriteFile(t *testing.T) {
	filename := "testfile.txt"

	err := os.WriteFile(filename, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}
	// Additional checks can be performed here
}
	