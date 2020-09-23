// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

const (
	exeFile = "tendon.exe"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Run

// Clean clean up after yourself
func Clean() {
	fmt.Println("Clean...")
	os.Remove(exeFile)
}

// Build build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(Clean)
	fmt.Println("Build...")
	cmd := exec.Command("go", "build", "-o", exeFile, ".")
	return cmd.Run()
}

// Run execute app
func Run() error {
	mg.Deps(Build)
	fmt.Println("Run...")
	cmd := exec.Command("./" + exeFile)
	return cmd.Run()
}
