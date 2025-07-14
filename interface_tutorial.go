package main

import (
    "fmt"
    "reflect"
    "strconv"
)

func main7() {
    // 1. What is interface{}?
    // interface{} is the empty interface - it has zero methods
    // Since every type implements at least zero methods,
    // interface{} can hold ANY value of ANY type

    fmt.Println("=== Understanding interface{} ===")

    // 2. Basic usage - storing different types
    fmt.Println("\n--- Storing Different Types ---")

    var anything interface{}

    anything = 42
    fmt.Printf("Storing int: %v (type: %T)\n", anything, anything)

    anything = "Hello"
    fmt.Printf("Storing string: %v (type: %T)\n", anything, anything)

    anything = []int{1, 2, 3}
    fmt.Printf("Storing slice: %v (type: %T)\n", anything, anything)

    anything = map[string]int{"age": 25}
    fmt.Printf("Storing map: %v (type: %T)\n", anything, anything)

    // 3. Type assertions - getting the original type back
    fmt.Println("\n--- Type Assertions ---")
    demonstrateTypeAssertions()

    // 4. Type switches - handling multiple types
    fmt.Println("\n--- Type Switches ---")
    demonstrateTypeSwitches()

    // 5. Common usage patterns
    fmt.Println("\n--- Common Usage Patterns ---")

    // Pattern 1: Functions that accept any type
    printAnything("Hello")
    printAnything(42)
    printAnything([]string{"a", "b", "c"})

    // Pattern 2: Collections of mixed types
    demonstrateMixedCollections()

    // Pattern 3: JSON-like data structures
    demonstrateJSONLike()

    // Pattern 4: Configuration and settings
    demonstrateConfiguration()

    // Pattern 5: Generic containers (before Go 1.18 generics)
    demonstrateGenericContainers()

    // 6. Working with reflection
    fmt.Println("\n--- Working with Reflection ---")
    demonstrateReflection()

    // 7. Best practices and pitfalls
    fmt.Println("\n--- Best Practices ---")
    demonstrateBestPractices()
}

func demonstrateTypeAssertions() {
    var value interface{} = "Hello, World!"

    // Safe type assertion with ok pattern
    if str, ok := value.(string); ok {
        fmt.Printf("Successfully got string: %s\n", str)
    } else {
        fmt.Println("Value is not a string")
    }

    // Unsafe type assertion (can panic)
    // str := value.(string) // This would work but could panic if wrong type

    // Trying to assert wrong type
    if num, ok := value.(int); ok {
        fmt.Printf("Got int: %d\n", num)
    } else {
        fmt.Println("Value is not an int")
    }
}

func demonstrateTypeSwitches() {
    values := []interface{}{
        42,
        "hello",
        3.14,
        []int{1, 2, 3},
        map[string]int{"count": 5},
        true,
        nil,
    }

    for i, value := range values {
        fmt.Printf("Value %d: ", i)

        switch v := value.(type) {
        case int:
            fmt.Printf("Integer: %d\n", v)
        case string:
            fmt.Printf("String: %q\n", v)
        case float64:
            fmt.Printf("Float: %.2f\n", v)
        case []int:
            fmt.Printf("Int slice with %d elements: %v\n", len(v), v)
        case map[string]int:
            fmt.Printf("String-to-int map: %v\n", v)
        case bool:
            fmt.Printf("Boolean: %t\n", v)
        case nil:
            fmt.Println("Nil value")
        default:
            fmt.Printf("Unknown type: %T with value %v\n", v, v)
        }
    }
}

// Pattern 1: Functions that accept any type
func printAnything(value interface{}) {
    fmt.Printf("Received: %v (type: %T)\n", value, value)
}

// Pattern 2: Collections of mixed types
func demonstrateMixedCollections() {
    fmt.Println("\n--- Mixed Collections ---")

    // Slice of mixed types
    items := []interface{}{
        "John",
        25,
        true,
        []string{"hobby1", "hobby2"},
    }

    fmt.Println("Mixed slice:")
    for i, item := range items {
        fmt.Printf("  [%d]: %v (%T)\n", i, item, item)
    }

    // Map with interface{} values
    person := map[string]interface{}{
        "name":    "Alice",
        "age":     30,
        "married": true,
        "skills":  []string{"Go", "Python", "JavaScript"},
    }

    fmt.Println("\nMixed map:")
    for key, value := range person {
        fmt.Printf("  %s: %v (%T)\n", key, value, value)
    }
}

// Pattern 3: JSON-like data structures
func demonstrateJSONLike() {
    fmt.Println("\n--- JSON-like Data ---")

    // Nested data structure similar to JSON
    data := map[string]interface{}{
        "users": []interface{}{
            map[string]interface{}{
                "id":   1,
                "name": "John",
                "profile": map[string]interface{}{
                    "email": "john@example.com",
                    "age":   25,
                },
            },
            map[string]interface{}{
                "id":   2,
                "name": "Jane",
                "profile": map[string]interface{}{
                    "email": "jane@example.com",
                    "age":   28,
                },
            },
        },
        "total": 2,
    }

    // Accessing nested data
    if users, ok := data["users"].([]interface{}); ok {
        fmt.Printf("Found %d users:\n", len(users))
        for i, user := range users {
            if userMap, ok := user.(map[string]interface{}); ok {
                name := userMap["name"]
                fmt.Printf("  User %d: %v\n", i+1, name)
            }
        }
    }
}

// Pattern 4: Configuration and settings
type Config struct {
    settings map[string]interface{}
}

func NewConfig() *Config {
    return &Config{
        settings: make(map[string]interface{}),
    }
}

func (c *Config) Set(key string, value interface{}) {
    c.settings[key] = value
}

func (c *Config) Get(key string) interface{} {
    return c.settings[key]
}

func (c *Config) GetString(key string, defaultValue string) string {
    if value, ok := c.settings[key].(string); ok {
        return value
    }
    return defaultValue
}

func (c *Config) GetInt(key string, defaultValue int) int {
    if value, ok := c.settings[key].(int); ok {
        return value
    }
    return defaultValue
}

func demonstrateConfiguration() {
    fmt.Println("\n--- Configuration System ---")

    config := NewConfig()

    // Set various types of configuration
    config.Set("host", "localhost")
    config.Set("port", 8080)
    config.Set("debug", true)
    config.Set("timeouts", []int{30, 60, 120})

    // Get configuration with type safety
    host := config.GetString("host", "127.0.0.1")
    port := config.GetInt("port", 3000)

    fmt.Printf("Server config: %s:%d\n", host, port)

    // Direct access (requires type assertion)
    if debug, ok := config.Get("debug").(bool); ok && debug {
        fmt.Println("Debug mode enabled")
    }
}

// Pattern 5: Generic containers (pre-generics era)
type Container struct {
    items []interface{}
}

func NewContainer() *Container {
    return &Container{
        items: make([]interface{}, 0),
    }
}

func (c *Container) Add(item interface{}) {
    c.items = append(c.items, item)
}

func (c *Container) Get(index int) interface{} {
    if index >= 0 && index < len(c.items) {
        return c.items[index]
    }
    return nil
}

func (c *Container) ForEach(fn func(interface{})) {
    for _, item := range c.items {
        fn(item)
    }
}

func demonstrateGenericContainers() {
    fmt.Println("\n--- Generic Container ---")

    container := NewContainer()

    // Add different types
    container.Add("Hello")
    container.Add(42)
    container.Add([]int{1, 2, 3})

    // Process all items
    container.ForEach(func(item interface{}) {
        fmt.Printf("Container item: %v (%T)\n", item, item)
    })
}

// Working with reflection
func demonstrateReflection() {
    values := []interface{}{
        42,
        "hello",
        []int{1, 2, 3},
        map[string]int{"key": 123},
    }

    for _, value := range values {
        analyzeValue(value)
    }
}

func analyzeValue(value interface{}) {
    v := reflect.ValueOf(value)
    t := reflect.TypeOf(value)

    fmt.Printf("Value: %v\n", value)
    fmt.Printf("  Type: %v\n", t)
    fmt.Printf("  Kind: %v\n", v.Kind())
    fmt.Printf("  Can be converted to string: %t\n", v.Type().ConvertibleTo(reflect.TypeOf("")))

    // Example: convert to string if possible
    if str := convertToString(value); str != "" {
        fmt.Printf("  As string: %q\n", str)
    }
    fmt.Println()
}

func convertToString(value interface{}) string {
    switch v := value.(type) {
    case string:
        return v
    case int:
        return strconv.Itoa(v)
    case float64:
        return strconv.FormatFloat(v, 'f', -1, 64)
    case bool:
        return strconv.FormatBool(v)
    default:
        return fmt.Sprintf("%v", v)
    }
}

func demonstrateBestPractices() {
    fmt.Println("Best Practices for interface{}:")
    fmt.Println("1. Use interface{} sparingly - prefer specific types when possible")
    fmt.Println("2. Always use the 'ok' pattern for type assertions")
    fmt.Println("3. Consider using type switches for handling multiple types")
    fmt.Println("4. Document what types your functions expect")
    fmt.Println("5. With Go 1.18+, consider using generics instead of interface{}")

    // Example: Better to use specific interface
    fmt.Println("\n--- Better Alternative: Specific Interface ---")

    // Instead of interface{}, define what you actually need
    var stringer interface {
        String() string
    }

    // This is more type-safe and self-documenting
    stringer = CustomType{value: "Hello"}
    fmt.Printf("Using specific interface: %s\n", stringer.String())
}

type CustomType struct {
    value string
}

func (c CustomType) String() string {
    return fmt.Sprintf("CustomType: %s", c.value)
}

// Additional utility functions for working with interface{}
func isNil(value interface{}) bool {
    return value == nil || reflect.ValueOf(value).IsNil()
}

func deepEqual(a, b interface{}) bool {
    return reflect.DeepEqual(a, b)
}

func clone(value interface{}) interface{} {
    // Simple clone using reflection (warning: not deep clone)
    v := reflect.ValueOf(value)
    if v.Kind() == reflect.Ptr {
        return reflect.Indirect(v).Interface()
    }
    return value
}
