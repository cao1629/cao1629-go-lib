package main

import "fmt"

type Info struct {
    name    string
    age     int
    address string
}

type BasicInfo struct {
    name string
    age  int
}

func GetBasicInfoSlice(infos []Info, start, end int) []BasicInfo {
    if start < 0 {
        start = 0
    }
    if end > len(infos) {
        end = len(infos)
    }
    if start >= end {
        return []BasicInfo{}
    }

    result := make([]BasicInfo, 0, end-start)
    for i := start; i < end; i++ {
        result = append(result, BasicInfo{
            name: infos[i].name,
            age:  infos[i].age,
        })
    }
    return result
}

func PrintBasicInfoSlice(basicInfos []BasicInfo) {
    fmt.Println("BasicInfo slice:")
    for i, info := range basicInfos {
        fmt.Printf("[%d] Name: %s, Age: %d\n", i, info.name, info.age)
    }
}

func main() {
    // Sample data
    infos := []Info{
        {name: "Alice Johnson", age: 25, address: "123 Main St"},
        {name: "Bob Smith", age: 30, address: "456 Oak Ave"},
        {name: "Charlie Brown", age: 35, address: "789 Pine Rd"},
        {name: "Diana Prince", age: 28, address: "321 Elm St"},
        {name: "Eve Wilson", age: 32, address: "654 Maple Ave"},
    }

    // Get a slice of BasicInfo from index 1 to 4
    basicInfos := GetBasicInfoSlice(infos, 1, 4)

    // Print the result
    PrintBasicInfoSlice(basicInfos)

    // Print original slice for comparison
    fmt.Println("\nOriginal Info slice:")
    for i, info := range infos {
        fmt.Printf("[%d] Name: %s, Age: %d, Address: %s\n", i, info.name, info.age, info.address)
    }

    fmt.Printf("%v\n", basicInfos)
}
