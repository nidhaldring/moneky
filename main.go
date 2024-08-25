package main

import (
	"log"
	"monkey/repl"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		log.Fatal("Not implemented yet!")
	}
	repl.StartRepl()
}
