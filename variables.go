package main

import (
	"flag"
)

var install bool
var version bool
var run bool
var clean bool
var configData ConfigData
var data Data

type Data struct {
	Type            string
	Path            string
	BuildBaseImage  string
	DeployBaseImage string
	AppName         string
}

type ConfigData struct {
	Options struct {
		Type string `json:"type"`
		Path string `json:"path"`
	} `json:"options"`
	DockerImage struct {
		BuildBaseImage  string `json:"buildBaseImage"`
		DeployBaseImage string `json:"deployBaseImage"`
	} `json:"dockerImage"`
	DockerCompose struct {
		AppName string `json:"appName"`
	} `json:"dockerCompose"`
}

func init() {
	flag.BoolVar(&install, "install", false, "Build a main structure")
	flag.BoolVar(&version, "version", false, "Outputs Smoothly version")
	flag.BoolVar(&run, "run", false, "Runs the solution locally with docker")
	flag.BoolVar(&clean, "clean", false, "Clean the install")
	data.Path = "test-app/"
	data.Type = "node"
	data.BuildBaseImage = "node:18-alpine"
	data.DeployBaseImage = "nginx:alpine"
	data.AppName = "your-app"
}
