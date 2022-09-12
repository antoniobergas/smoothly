package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var build bool
var version bool
var start bool
var run bool

func init() {
	flag.BoolVar(&build, "build", false, "Build a main structure")
	flag.BoolVar(&version, "version", false, "Outputs Smoothly version")
	flag.BoolVar(&run, "run", false, "Runs the solution locally with docker")
	flag.BoolVar(&start, "start", false, "Runs the solution locally with docker")
}

func runBuild (){
	fmt.Println("Initializing structure...")
	file, err := os.Create("Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created successfully")
	defer file.Close()
}

func runVersion(){
	fmt.Println("1.0.0")
}

func runSolution(){
	cmd := exec.Command("docker", "build", "-t", "your-app", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	flag.Parse()

	if build {
		runBuild()
	}

	if version {
		runVersion()
	}

	if start {
		runBuild()
		runSolution()
	}

	if run {
		runSolution()
	}
}
