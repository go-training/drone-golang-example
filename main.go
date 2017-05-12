package main

import "fmt"

func HelloWorld() string {
	return "Hello World!!"
}

func main() {
	hello := HelloWorld()
	fmt.Println(hello)
}
