package main

import (
    "io"
    "log"
    "net"
    "os"
    "testing"
)

func server() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        // exit the program if we can't listen on the port
        // cases: port already in use, no permission, etc.
        log.Fatal(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err) // e.g., connection aborted
            continue
        }
        go handleConn3(conn) // handle one connection at a time
    }
}

func handleConn3(c net.Conn) {
    defer c.Close()

    io.Copy(c, os.Stdin)
}

func TestAA(t *testing.T) {
    server()
}
