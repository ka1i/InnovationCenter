package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ka1i/InnovationCenter/pkg/usage"
	"github.com/ka1i/InnovationCenter/pkg/version"
)

func flags() {
	if len(os.Args) > 1 {
		var argv = os.Args[1:]
		switch argv[0] {
		case "-h", "--help", "help":
			usage.Usage()
		case "-v", "--version", "version":
			fmt.Printf("%s version: %s\n", os.Args[0], version.VERSION)
		default:
			log.Println("please check usage")
		}
		os.Exit(0)
	}
}
