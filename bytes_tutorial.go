package main

import (
    "fmt"
)

func main6() {
    // 1. What is []byte?
    // []byte is a slice of bytes (uint8 values from 0-255)
    // It's commonly used for:
    // - Binary data
    // - File I/O
    // - Network communication
    // - String manipulation

    fmt.Println("=== Understanding []byte ===")

    // 2. Creating []byte in different ways
    fmt.Println("\n--- Creating []byte ---")

    // Method 1: From string literal
    data1 := []byte("Hello, World!")
    fmt.Printf("From string: %v\n", data1)
    fmt.Printf("As string: %s\n", string(data1))

    // Method 2: Using make()
    data2 := make([]byte, 5) // Creates slice with 5 zero bytes
    fmt.Printf("Using make(5): %v\n", data2)

    // Method 3: Byte literal
    data3 := []byte{72, 101, 108, 108, 111} // ASCII for "Hello"
    fmt.Printf("Byte literal: %v\n", data3)
    fmt.Printf("As string: %s\n", string(data3))

    // Method 4: Using make() with capacity
    data4 := make([]byte, 3, 10) // length=3, capacity=10
    fmt.Printf("make(3,10): %v (len=%d, cap=%d)\n", data4, len(data4), cap(data4))

    // 3. Common operations with []byte
    fmt.Println("\n--- Common Operations ---")

    // Appending bytes
    message := []byte("Go")
    message = append(message, ' ')
    message = append(message, []byte("is awesome!")...)
    fmt.Printf("After append: %s\n", string(message))

    // Indexing and slicing
    fmt.Printf("First byte: %d ('%c')\n", message[0], message[0])
    fmt.Printf("Slice [0:2]: %s\n", string(message[0:2]))

    // Modifying bytes
    message[0] = 'g' // Change 'G' to 'g'
    fmt.Printf("After modification: %s\n", string(message))

    // 4. Converting between string and []byte
    fmt.Println("\n--- String <-> []byte Conversion ---")

    text := "Hello, 世界"
    bytes := []byte(text)
    backToString := string(bytes)

    fmt.Printf("Original string: %s\n", text)
    fmt.Printf("As []byte: %v\n", bytes)
    fmt.Printf("Back to string: %s\n", backToString)
    fmt.Printf("Length in bytes: %d\n", len(bytes))
    fmt.Printf("Length in runes: %d\n", len([]rune(text)))

    // 5. Working with binary data
    fmt.Println("\n--- Binary Data ---")

    // Creating binary data
    binary := make([]byte, 4)
    binary[0] = 0xFF // 255
    binary[1] = 0x00 // 0
    binary[2] = 0xAB // 171
    binary[3] = 0xCD // 205

    fmt.Printf("Binary data: %v\n", binary)
    fmt.Printf("As hex: %x\n", binary)
    fmt.Printf("As hex with spaces: % x\n", binary)

    // 6. Practical examples
    fmt.Println("\n--- Practical Examples ---")

    // Example 1: Simple text processing
    processText()

    // Example 2: Building a buffer
    buildBuffer()

    // Example 3: Parsing binary data
    parseBinary()
}

func processText() {
    fmt.Println("\nExample 1: Text Processing")

    text := "hello world"
    data := []byte(text)

    // Convert to uppercase by modifying bytes
    for i := 0; i < len(data); i++ {
        if data[i] >= 'a' && data[i] <= 'z' {
            data[i] = data[i] - 32 // Convert to uppercase
        }
    }

    fmt.Printf("Original: %s\n", text)
    fmt.Printf("Uppercase: %s\n", string(data))
}

func buildBuffer() {
    fmt.Println("\nExample 2: Building a Buffer")

    var buffer []byte

    // Build a CSV-like format
    buffer = append(buffer, []byte("Name,Age,City\n")...)
    buffer = append(buffer, []byte("Alice,30,NYC\n")...)
    buffer = append(buffer, []byte("Bob,25,LA\n")...)

    fmt.Printf("CSV Buffer:\n%s", string(buffer))
    fmt.Printf("Buffer size: %d bytes\n", len(buffer))
}

func parseBinary() {
    fmt.Println("\nExample 3: Parsing Binary Data")

    // Simulate a simple binary format: [length][data]
    message := "Go rocks!"

    // Create binary packet: 1 byte length + message
    packet := make([]byte, 1+len(message))
    packet[0] = byte(len(message))    // Store length in first byte
    copy(packet[1:], []byte(message)) // Copy message after length

    fmt.Printf("Binary packet: %v\n", packet)
    fmt.Printf("Packet hex: %x\n", packet)

    // Parse the packet
    if len(packet) > 0 {
        msgLength := int(packet[0])
        if len(packet) >= 1+msgLength {
            extractedMessage := string(packet[1 : 1+msgLength])
            fmt.Printf("Extracted message: %s\n", extractedMessage)
        }
    }
}

// Bonus: Useful helper functions for []byte
func demonstrateHelpers() {
    fmt.Println("\n=== Useful Helper Functions ===")

    data := []byte("Hello, World!")

    // Check if contains
    if containsByte(data, 'W') {
        fmt.Println("Contains 'W'")
    }

    // Find index
    index := findByte(data, ',')
    fmt.Printf("Comma at index: %d\n", index)

    // Replace bytes
    replaced := replaceByte(data, 'o', '0')
    fmt.Printf("After replacing 'o' with '0': %s\n", string(replaced))
}

func containsByte(data []byte, target byte) bool {
    for _, b := range data {
        if b == target {
            return true
        }
    }
    return false
}

func findByte(data []byte, target byte) int {
    for i, b := range data {
        if b == target {
            return i
        }
    }
    return -1
}

func replaceByte(data []byte, old, new byte) []byte {
    result := make([]byte, len(data))
    for i, b := range data {
        if b == old {
            result[i] = new
        } else {
            result[i] = b
        }
    }
    return result
}
