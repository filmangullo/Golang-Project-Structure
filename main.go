main.go

package main

import "fmt"

func init() {
	// Call the Database function from the initializers package
	initializers.Database()
}

func main() {
    fmt.Println("Hello")
}
