package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Hello World!!" {
		t.Error("Testing error")
	}
}
