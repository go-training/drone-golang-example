package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Hello GitHub Actions, traefik workshop!" {
		t.Error("Testing error")
	}
}
