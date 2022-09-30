package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func getData() {
	file, _ := os.ReadFile("test-app/smoothly.json")
	_ = json.Unmarshal([]byte(file), &configData)
}

func createFromTemplate[T any](path string, fileName string, data T) {

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

// func createFile(content string) {

// 	f, err := os.Create("test-app/smoothly.json")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	_, err2 := f.WriteString(content)

// 	if err2 != nil {
// 		log.Fatal(err2)
// 	}
// }

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
