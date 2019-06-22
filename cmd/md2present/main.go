package main

import "fmt"

var (
	Branch   string
	Revision string
	Version  string
	Created  string
)

func main() {
  fmt.Println("hello", Branch, Revision, Version, Created)
}
