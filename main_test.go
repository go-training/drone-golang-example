package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Hello World, drone workshop!" {
		t.Error("Testing error")
	}
}
