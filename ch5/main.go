package main

import (
	"learning/ch5-gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
