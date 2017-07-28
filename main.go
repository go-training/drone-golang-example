package main

import "fmt"

// HelloWorld for hello world
func HelloWorld() string {
	return "Hello World, drone workshop!!"
}

func main() {
	hello := HelloWorld()
	fmt.Println(hello)
}
