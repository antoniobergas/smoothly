package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

var install bool
var version bool
var start bool
var run bool

type Dockerfile struct {
	Message string
}

type DockerCompose struct {
	Message string
}

func init() {
	flag.BoolVar(&install, "install", false, "Build a main structure")
	flag.BoolVar(&version, "version", false, "Outputs Smoothly version")
	flag.BoolVar(&run, "run", false, "Runs the solution locally with docker")
	flag.BoolVar(&start, "start", false, "Runs the solution locally with docker")
}

func createFromTemplate[T any](fileName string, data T) {

	var content bytes.Buffer

	tmpl, err := template.ParseFiles(fileName + ".template")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(&content, &data)
	if err != nil {
		log.Fatalln(err)
	}

	os.WriteFile(fileName, content.Bytes(), 0644)
}

func runInstall() {
	fmt.Println("Initializing structure...")
	dockerfileData := Dockerfile{Message: "Hello World"}
	dockercomposeData := DockerCompose{Message: "Hello World"}
	createFromTemplate("Dockerfile", dockerfileData)
	createFromTemplate("docker-compose.yml", dockercomposeData)
}

func runVersion() {
	fmt.Println("1.0.0")
}

func runSolution() {
	cmd := exec.Command("docker", "build", "-t", "your-app", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	flag.Parse()

	if install {
		runInstall()
	}

	if version {
		runVersion()
	}

	if start {
		runInstall()
		runSolution()
	}

	if run {
		runSolution()
	}
}
