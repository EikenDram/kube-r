package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

// api structure
type ApiStruct struct {
	Path string `json:"path"`
	File string `json:"file"`
}

// config.json structure
type ConfigJson struct {
	Api []ApiStruct `json:"api"`
}

// configuration
var config ConfigJson

// template data
var templateData TemplateMap

// env variables
var user string
var pass string

// template map structure
type TemplateMap struct {
	Api []ApiStruct
}

func loadConfig() {
	// read env variables
	user = os.Getenv("SERVER_USER")
	pass = os.Getenv("SERVER_PASS")

	// read config.json
	file, err := os.ReadFile("./config.json")
	check(err)
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	//create map
	templateData = TemplateMap{config.Api}
}

// executes template tmpl from filenames to file fout using loaded values
func buildTemplate(fileOut string, templateName string, filenames ...string) {
	os.Remove(fileOut)

	t, err := template.New("").
		Option("missingkey=zero").
		ParseFiles(filenames...)
	check(err)

	f, err := os.OpenFile(fileOut, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)

	err = t.ExecuteTemplate(f, templateName, templateData)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func executeCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}

	defer stdOut.Close()

	scanner := bufio.NewScanner(stdOut)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(os.Stderr, "Error creating StderrPipe for Cmd", err)
	}

	defer stdErr.Close()

	stdErrScanner := bufio.NewScanner(stdErr)
	go func() {
		for stdErrScanner.Scan() {
			fmt.Printf("%s\n", stdErrScanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(os.Stderr, "Error starting Cmd", err)
	}

	err = cmd.Wait()
	// go generate command will fail when no generate command find.
	if err != nil {
		if err.Error() != "exit status 1" {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
}
