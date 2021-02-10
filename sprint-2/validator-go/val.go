package main

import (
	"360-richard-goluszka/sprint-2/validator-go/directorychk"
	"360-richard-goluszka/sprint-2/validator-go/licensechk"
	"360-richard-goluszka/sprint-2/validator-go/linefmtchk"
	"360-richard-goluszka/sprint-2/validator-go/utf8chk"
	"fmt"
	"os"
	"strings"
)

const borderChar = `=`
const borderLen = 80

type validator interface {
	Validate() bool
	GetMsg() string
	GetIssues() string
	GetIssueCt() int
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

func dispMsgs(borderLength int, messages ...string) {
	for _, msg := range messages {
		dispMsg(msg, borderLength)
	}
}

//handle command-line arguments and then exit
func checkArgs(args []string) bool {
	if strings.EqualFold(args[0], `help`) {
		dispBorder(borderChar, borderLen)
		dispMsgs(borderLen, `val.go -- validates code compliance with the following standards:`,
			`1. Only files required to compile/execute plus README and LICENSE are included`,
			`4. A LICENSE file included that is MIT, GNU GPL, or all rights reserved`,
			`6. Tabs are used for indenting and line feeds (\n) mark the end of lines`,
			`13. All files are UTF-8 compatible text files`, ``,
			"\u2018val\u2019 reports the number of errors and PASS/FAIL status",
			"\u2018val detail\u2019 shows files checked and error locations",
			"\u2018val help\u2019 displays this information again")
		dispBorder(borderChar, borderLen)
		os.Exit(0)
	} else if strings.EqualFold(args[0], `detail`) {
		dispBorder(borderChar, borderLen)
		dispMsg(`Detail Mode: On`, borderLen)
		return true
	} else { //handle invalid options
		dispBorder(borderChar, borderLen)
		dispMsgs(borderLen, "Invalid argument: \u2018"+args[0]+"\u2019",
			"use \u2018val help\u2019 to find valid arguments")
		dispBorder(borderChar, borderLen)
		os.Exit(0)
	}
	return false
}

func main() {
	//parse command-line arguments
	args := os.Args[1:]
	detailMode := false
	if len(args) > 0 {
		detailMode = checkArgs(args)
	}

	//create a slice to hold the validation structs
	valUnits := []validator{
		&directorychk.DirChecker{Path: `.`},
		&licensechk.LicenseChecker{Path: `.`},
		&linefmtchk.LineFmtChecker{Path: `.`},
		&utf8chk.UTF8Checker{Path: `.`},
	}

	//create a slice to hold coding standard labels (for UI)
	valLabels := []string{
		`Directory Contents Check`,
		`License Contents Check`,
		`Line Feed and Tabs Check`,
		`UTF8 Compatability Check`,
	}

	//create a slice to hold coding standard descriptions (for UI)
	valDescs := []string{
		`Ensures files are LICENSE, README.md, or end with .go or .mod`,
		`LICENSE file mentions MIT, GNU, or all rights reserved`,
		`File lines do not end in \n and code is not space-indented`,
		`File lines contain only valid UTF-8 characters`,
	}

	//run each validator and display output polymorphically
	status := ``
	dispBorder(borderChar, borderLen)
	for index, unit := range valUnits {
		if unit.Validate() {
			status = `Status: PASS`
		} else {
			status = `Status: FAIL`
		}
		status += fmt.Sprint("\tIssue Count: ", unit.GetIssueCt())
		dispMsgs(borderLen, valLabels[index], valDescs[index], ``)
		if detailMode { //output for `detail` argument
			fmt.Println(unit.GetMsg())
			fmt.Println(unit.GetIssues())
		}
		dispMsg(status, borderLen)
		dispBorder(borderChar, borderLen)
	}
}
