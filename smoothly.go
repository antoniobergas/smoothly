package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
var configData ConfigData

type ConfigData struct {
	Options struct {
		Path string `json:"path"`
	} `json:"options"`
	DockerImage struct {
		Version      string `json:"version"`
	} `json:"dockerImage"`
	DockerCompose struct {
		AppName      string `json:"appName"`
	} `json:"dockerCompose"`
}

func init() {
	flag.BoolVar(&install, "install", false, "Build a main structure")
	flag.BoolVar(&version, "version", false, "Outputs Smoothly version")
	flag.BoolVar(&run, "run", false, "Runs the solution locally with docker")
	flag.BoolVar(&start, "start", false, "Runs the solution locally with docker")
	flag.BoolVar(&clean, "clean", false, "Clean the install")
	file, _ := ioutil.ReadFile("test-app/smoothly.json")
	_ = json.Unmarshal([]byte(file), &configData)

}

func createFromTemplate[T any](path string, fileName string,  data T) {

	var content bytes.Buffer
	tmpl, err := template.ParseFiles("templates/" + fileName + ".template")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(&content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(path+fileName, content.Bytes(), 0644)
}

func outputCmd(cmd *exec.Cmd) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))
}

func removeFile(file string) {
	rm := os.Remove(file)
	if rm != nil {
		log.Fatal(rm)
	}
}

func runSolution() {
	fmt.Println("Running solution...")
	dockerBuild := exec.Command("docker", "build", "-t", "your-app-image", "test-app/.")
	outputCmd(dockerBuild)
	dockerCompose := exec.Command("docker", "compose", "-f", "test-app/docker-compose.yml", "up", "-d")
	outputCmd(dockerCompose)
	fmt.Println("Solution running!")
}

func runClean() {
	fmt.Println("Cleaning structure...")
	removeFile(configData.Options.Path + "Dockerfile")
	removeFile(configData.Options.Path + "docker-compose.yml")
	fmt.Println("Structure cleaned!")
}

func runInstall() {
	fmt.Println("Initializing structure...")
	createFromTemplate(configData.Options.Path, "Dockerfile", configData.DockerImage)
	createFromTemplate(configData.Options.Path, "docker-compose.yml", configData.DockerCompose)
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
