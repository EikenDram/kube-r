package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Starting KubeR server...")

	// load configuration
	fmt.Println("Loading configuration...")
	loadConfig()

	// configure nginx
	fmt.Println("Configuring nginx htpasswd...")
	htpasswd()

	// build templates
	fmt.Println("Building templates..")
	buildTemplate("packages.R", "packages.R", "templates/packages.R")
	buildTemplate("plumber.R", "plumber.R", "templates/plumber.R")

	// install packages?
	fmt.Println("Installing additional packages..")
	installPackages()

	// start plumber api
	fmt.Println("Staring plumber API...")
	startApi()
}

// generate htpasswd file
func htpasswd() {
	os.Remove("/etc/nginx/.htpasswd")

	app := "htpasswd"

	arg0 := "-bc"
	arg1 := "/etc/nginx/.htpasswd"
	arg2 := user
	arg3 := pass

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

// run packages.R script
func installPackages() {
	// this will be run on each restart i think
	// 2D: might leave as-is, think R checks that package is already installed and wont install again
	app := "Rscript"

	arg0 := "packages.R"

	executeCmd(app, arg0)
}

// start plumber API
func startApi() {
	app := "Rscript"

	arg0 := "api.R"

	executeCmd(app, arg0)
}
