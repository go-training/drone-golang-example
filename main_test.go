package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Hello World, golang workshop!" {
		t.Error("Testing error")
	}
}
