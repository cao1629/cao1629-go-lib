package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
    "testing"
)

func TestA(t *testing.T) {
    buf := bytes.NewBufferString("hello")

    // Read all data
    data := make([]byte, 10)

    // try to read 10 bytes from this buffer because len(data) is 10
    n, err := buf.Read(data)
    fmt.Printf("Read %d bytes: %s, err: %v\n", n, string(data[:n]), err)
    // Read 5 bytes, err is nil. This buffer is now drained.

    // Try to read from drained buffer
    n, err = buf.Read(data)
    fmt.Printf("Read %d bytes, err: %v\n", n, err)
    // Output: Read 0 bytes, err: EOF

}

func TestB(t *testing.T) {
    var buf bytes.Buffer
    n, err := buf.WriteTo(os.Stdout)
    fmt.Printf("Write %d bytes: %s, err: %v\n", n, buf.String(), err)
}

func TestC(t *testing.T) {
    // Create a string and we will read data from this string later
    r := strings.NewReader("some io.Reader stream to be read\n")

    // os.Stdout is a File, which is a Writer
    // io.Copy   从Reader读数据到Writer 一直读到EOF
    //
    // 一直读到EOF什么意思?
    // file: 读到文件结束
    // tcp connection: 读到connection断了
    if _, err := io.Copy(os.Stdout, r); err != nil {
        log.Fatal(err)
    }

    if _, err := io.Copy(os.Stdout, r); err != nil {
        log.Fatal(err)
    }

}

func TestD(t *testing.T) {
    b := bytes.NewBufferString("ABC\n")
    b.WriteTo(os.Stdout)
    _, err := b.WriteTo(os.Stdout) // 这个会报错，因为buffer已经读完了
    if err != nil {
        fmt.Printf("Error writing to stdout: %v\n", err)
    } else {
        fmt.Println("WriteTo successful")
    }
}
