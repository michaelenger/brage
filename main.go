package main

import (
	"github.com/michaelenger/brage/cmd"
	"log"
)

func main() {
	logger := log.Default()
	logger.SetFlags(0) // disable the date/time

	cmd.Execute()
}
