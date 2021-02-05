package main

import (
	"360-richard-goluszka/sprint-2/validator-go/directorychk"
	"360-richard-goluszka/sprint-2/validator-go/licensechk"
	"fmt"
	"os"
	"strings"
)

const borderLength = 80

type validator interface {
	Validate() bool
	GetMsg() string
}

//print a border of a given string with a given length
func dispBorder(borderChar string, borderLength int) {
	border := strings.Repeat(borderChar, borderLength)
	fmt.Println(border)
}

//print centered messages based on given borderLength
func dispMsg(msg string, borderLength int) {
	padding := borderLength - len(msg)
	dispMsg := strings.Repeat(` `, padding/2) + msg
	fmt.Println(dispMsg)
}

//handle command-line arguments passed to executable during call
func checkArgs(args []string) {
	if strings.ToLower(args[0]) == `help` { //handle cli 'help' argument and exit
		dispBorder(`=`, borderLength)
		dispMsg(`val.go -- written to automate validation of coding standards`,
			borderLength)
		dispMsg("use \u2018val help\u2019 to display this information again", borderLength)
		dispBorder(`=`, borderLength)
		os.Exit(0)
	} else { //handle other cli arguments and exit
		dispBorder(`=`, borderLength)
		dispMsg("Invalid argument: \u2018"+args[0]+"\u2019", borderLength)
		dispMsg("use \u2018val help\u2019 to find valid arguments", borderLength)
		dispBorder(`=`, borderLength)
		os.Exit(0)
	}
}

func dispValidation(valUnits ...validator) {
	for _, unit := range valUnits {
		fmt.Println(unit.Validate())
		fmt.Println(unit.GetMsg())
	}
}

func main() {
	//parse any command-line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		checkArgs(args)
	}

	//create and call each validation struct polymorphically
	dc := directorychk.DirChecker{Path: `.`}
	lc := licensechk.LicenseChecker{Path: `.`}
	dispValidation(&dc, &lc)

}
