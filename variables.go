package main

import (
	"flag"
)

var install bool
var version bool
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
	flag.BoolVar(&clean, "clean", false, "Clean the install")
}