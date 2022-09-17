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
var clean bool
var dockerImageData DockerImage
var dockerComposeData DockerCompose

type DockerImage struct {
	Version string
}

type DockerCompose struct {
	AppName string
}

func init() {
	flag.BoolVar(&install, "install", false, "Build a main structure")
	flag.BoolVar(&version, "version", false, "Outputs Smoothly version")
	flag.BoolVar(&run, "run", false, "Runs the solution locally with docker")
	flag.BoolVar(&start, "start", false, "Runs the solution locally with docker")
	flag.BoolVar(&clean, "clean", false, "Clean the install")
	dockerImageData = DockerImage{Version: "18"}
	dockerComposeData = DockerCompose{AppName: "your-app"}
}

func createFromTemplate[T any](fileName string, data T) {

	var content bytes.Buffer

	tmpl, err := template.ParseFiles("templates/" + fileName + ".template")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(&content, &data)
	if err != nil {
		log.Fatalln(err)
	}

	os.WriteFile("test-app/" + fileName, content.Bytes(), 0644)
}

func outputCmd(cmd *exec.Cmd){
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))
}

func removeFile(fileName string){
	rm := os.Remove(fileName)
    if rm != nil {
        log.Fatal(rm)
    }
}

func runSolution() {
	fmt.Println("Running solution...")
	dockerBuild := exec.Command("docker", "build", "-t", "your-app-image", "test-app/.")
	outputCmd(dockerBuild)
	dockerCompose := exec.Command("docker", "compose", "-f", "test-app/docker-compose.yml","up", "-d")
	outputCmd(dockerCompose)
	fmt.Println("Solution running!")
}

func runClean() {
	fmt.Println("Cleaning structure...")
    removeFile("test-app/Dockerfile")
	removeFile("test-app/docker-compose.yml")
	fmt.Println("Structure cleaned!")
}

func runInstall() {
	fmt.Println("Initializing structure...")
	createFromTemplate("Dockerfile", &dockerImageData)
	createFromTemplate("docker-compose.yml", &dockerComposeData)
	fmt.Println("Structure created!")
}

func runVersion() {
	fmt.Println("smoothly-1.0.0")
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

	if clean {
		runClean()
	}
}
