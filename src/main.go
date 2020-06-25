package main

import (
	"log"

	"github.com/lachee/noodle"
)

//app is the application
var app = &SpagApplication{}

//var app = &RotatingCubeApp{}

func main() {

	//This is blocking. Everything after this happens in the application
	// It will return an exit code
	exitCode := noodle.Run(app)
	log.Println("Exited with code", exitCode)
}
