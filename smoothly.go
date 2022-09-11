package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {

	init := flag.Bool("init", false, "Initialize a main structure")

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

}
