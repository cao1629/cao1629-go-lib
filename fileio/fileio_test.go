package fileio

import (
    "os"
    "testing"
)

func TestWriteFile(t *testing.T) {
    filename := "testfile.txt"

    err := os.WriteFile(filename, []byte("Hello, World!"), 0644)
    if err != nil {
        t.Fatalf("Failed to write file: %v", err)
    }
    // Additional checks can be performed here
}

func TestNewFile(t *testing.T) {
    file, err := os.Create("1.txt")
    if err != nil {
        t.Fatalf("Failed to create file: %v", err)
    }

    defer file.Close()
    _, err = file.WriteString("Hello, World!")
    if err != nil {
        return
    }
}
