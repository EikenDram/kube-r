package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"gopkg.in/yaml.v2"
)

// configuration
var config ConfigYaml

// template data
var templateData TemplateMap

// env variables
var user string
var pass string

// load configuration from yaml and env
func loadConfig() {
	// read env variables
	user = os.Getenv("SERVER_USER")
	pass = os.Getenv("SERVER_PASS")

	// read config.json
	file, err := os.ReadFile("./config.yaml")
	check(err)
	err = yaml.Unmarshal(file, &config)
	check(err)

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

// check for error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// execute os command
func executeCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating StdoutPipe for Cmd")
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
		fmt.Fprintf(os.Stderr, "Error creating StderrPipe for Cmd")
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
		fmt.Fprintf(os.Stderr, "Error starting Cmd")
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
