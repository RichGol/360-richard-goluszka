package main

import (
	"fmt"
	"git/360-richard-goluszka/sprint-2/validator-go/dircheck"
	"os"
	"strings"
)

const borderLength = 80

//print a border of a given string with a given length
func displayBorder(borderChar string, borderLength int) {
	border := strings.Repeat(borderChar, borderLength)
	fmt.Println(border)
}

//print centered messages based on given borderLength
func displayMessage(msg string, borderLength int) {

	//center and display msg in a nice UI
	padding := borderLength - len(msg)
	dispMsg := strings.Repeat(` `, padding/2) + msg
	fmt.Println(dispMsg)
}

//handle command-line arguments passed to executable during call
func checkArgs(args []string) {
	if strings.ToLower(args[0]) == `help` { //handle cli 'help' argument and exit
		displayBorder(`=`, borderLength)
		displayMessage(`val.go -- written to automate validation of coding standards`,
			borderLength)
		displayMessage("use \u2018val help\u2019 to display this information again", borderLength)
		displayBorder(`=`, borderLength)
		os.Exit(0)
	} else { //handle other cli arguments and exit
		displayBorder(`=`, borderLength)
		displayMessage("Invalid argument: \u2018"+args[0]+"\u2019", borderLength)
		displayMessage("use \u2018val help\u2019 to find valid arguments", borderLength)
		displayBorder(`=`, borderLength)
		os.Exit(0)
	}
}

func main() {
	//parse any command-line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		checkArgs(args)
	}

	//validate standards
	dc := dircheck.DirChecker{Path: `.`, Msg: ``}
	fmt.Println(dc.Validate())
}
