package main

import (
	"360-richard-goluszka/sprint-2/validator-go/directorychk"
	"360-richard-goluszka/sprint-2/validator-go/licensechk"
	"360-richard-goluszka/sprint-2/validator-go/linefmtchk"
	"fmt"
	"os"
	"strings"
)

const borderChar = `=`
const borderLen = 80

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
		dispBorder(borderChar, borderLen)
		dispMsg(`val.go -- written to automate validation of coding standards`,
			borderLen)
		dispMsg("use \u2018val help\u2019 to display this information again", borderLen)
		dispBorder(borderChar, borderLen)
		os.Exit(0)
	} else { //handle other cli arguments and exit
		dispBorder(borderChar, borderLen)
		dispMsg("Invalid argument: \u2018"+args[0]+"\u2019", borderLen)
		dispMsg("use \u2018val help\u2019 to find valid arguments", borderLen)
		dispBorder(borderChar, borderLen)
		os.Exit(0)
	}
}

func main() {
	//parse any command-line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		checkArgs(args)
	}

	//create & associate each validator struct with a descriptive string
	valUnitMap := map[string]validator{
		`Directory Contents Check`: &directorychk.DirChecker{Path: `.`},
		`License Contents Check`:   &licensechk.LicenseChecker{Path: `.`},
		`Line Feed Endings Check`:  &linefmtchk.LineFmtChecker{Path: `.`},
	}

	//run and display the output of each validator struct
	status := ``
	dispBorder(borderChar, borderLen)
	for label, unit := range valUnitMap {
		if unit.Validate() {
			status = `Status: PASS`
		} else {
			status = `Status: FAIL`
		}
		dispMsg(label, borderLen)
		fmt.Println(unit.GetMsg())
		dispMsg(status, borderLen)
		dispBorder(borderChar, borderLen)
	}
}
