package main

import (
	"os/exec"
	"testing"
)

func TestCreateFromTemplate(t *testing.T) {
	type TestData struct {
		TestField string
	}
	var testData TestData
	testData.TestField = "Test data"
	createFromTemplate("testing/", "testfile", testData)

}

func TestOutputCmd(t *testing.T) {
	var cmdTest = exec.Command("echo", "cmd test")
	outputCmd(cmdTest)
}

func TestRemoveFile(t *testing.T) {
	removeFile("testing/testfile")
}