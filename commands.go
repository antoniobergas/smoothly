package main

import (
	"fmt"
	"os/exec"
)

func runInstall() {
	fmt.Println("Initializing structure...")
	createFile(
		`{
		"options": {
		  "path": "test-app/"
		},
		"dockerImage":{
		  "version": "18"
		},
		"dockerCompose":{
		  "appName": "your-app"
		}
	  }`)
	getData()
	createFromTemplate(configData.Options.Path, "Dockerfile", configData.DockerImage)
	createFromTemplate(configData.Options.Path, "docker-compose.yml", configData.DockerCompose)
	fmt.Println("Structure created!")
}

func runSolution() {
	getData()
	fmt.Println("Running solution...")
	dockerBuild := exec.Command("docker", "build", "-t", "your-app-image", configData.Options.Path + ".")
	outputCmd(dockerBuild)
	dockerCompose := exec.Command("docker", "compose", "-f", configData.Options.Path + "docker-compose.yml", "up", "-d")
	outputCmd(dockerCompose)
	fmt.Println("Solution running!")
}

func runClean() {
	getData()
	fmt.Println("Cleaning structure...")
	removeFile(configData.Options.Path + "Dockerfile")
	removeFile(configData.Options.Path + "docker-compose.yml")
	removeFile(configData.Options.Path + "smoothly.json")
	fmt.Println("Structure cleaned!")
}

func runVersion() {
	fmt.Println("smoothly-1.0.0")
}
