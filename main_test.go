package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Trigger the GitHub Actions, traefik workshop!" {
		t.Error("Testing error")
	}
}
