package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Hello World, traefik workshop!" {
		t.Error("Testing error")
	}
}
