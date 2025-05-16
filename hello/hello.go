package main

// import "fmt"
// import "rsc.io/quote"
// import "runtime"

import (
	"fmt"
	"runtime"
)


// Package-level variables
var appName string = "MyGoApp"
var version float64 = 1.0

var a [5]int // An array of 5 integers


func main() {
	fmt.Println("App Name:", appName, "Version:", version)

	os := runtime.GOOS

	fmt.Println("Operating System:", os)

	switch os {
	case "darwin":
		fmt.Println("Looks like macos")
	case "linux":
		fmt.Println("Looks like linux")
	case "windows":
		fmt.Println("Looks like windows")
	default:
		fmt.Printf("Uknown operating system %s\n", os)
	}


	// fmt.Println("Hello AMAL")
	// fmt.Println("Kata Kata Hari Ini : " + quote.Go())
	// name := "AMAL"
	// array := [...]float64{7.0, 8.5, 9.1}
	// price := Sum(&array) 
	// fmt.Printf("Name: %s, Price: %.2f\n", name, price)


	b := [3]string{"apple", "banana", "cherry"}

	fmt.Print(b[2])
	fmt.Println()
	
	var p Person
	person := Person{Name: "Charlie", Age: 35, Address: "123 Main St"}
	fmt.Println(person)
	fmt.Println(p)
}

func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

type Person struct {
    Name    string
    Age     int
    Address string
}
