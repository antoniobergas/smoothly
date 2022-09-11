package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	init := flag.Bool("init", false, "Initialize a main structure")
	version := flag.Bool("version", false, "Outputs Smoothly version")

	flag.Parse()

	if *init {
		fmt.Println("Initializing structure...")
		file, err := os.Create("Dockerfile")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File created successfully")
		defer file.Close()
	}

	if *version {
		fmt.Println("1.0.0")
	}

}
