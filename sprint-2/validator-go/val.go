package main

import (
	"360-richard-goluszka/sprint-2/validator-go/directorychk"
	"360-richard-goluszka/sprint-2/validator-go/licensechk"
	"360-richard-goluszka/sprint-2/validator-go/linefmtchk"
	"360-richard-goluszka/sprint-2/validator-go/utf8chk"
	"bufio"
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
		dispMsgs(borderLen, `val.go -- validates code compliance with standards 1, 4, 6, and 13`,
			"\u2018val\u2019 reports the number of errors and PASS/FAIL status",
			"\u2018val detail\u2019 shows the files checked and any error locations",
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

	//Get path from user
	fmt.Print(`Enter path to project: `)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	path := scanner.Text()

	//create a slice to hold coding standard structs
	valUnits := []validator{
		&directorychk.DirChecker{Path: path},
		&licensechk.LicenseChecker{Path: path},
		&linefmtchk.LineFmtChecker{Path: path},
		&utf8chk.UTF8Checker{Path: path},
	}

	//create a slice to hold coding standard labels (for UI)
	valLabels := []string{
		`Directory Contents Check`,
		`License Contents Check`,
		`Line Feed and Tabs Check`,
		`UTF8 Compatibility Check`,
	}

	//create a slice to hold coding standard descriptions (for UI)
	valDescs := []string{
		`Only files required to compile/execute, README.md, and LICENSE are included`,
		`The LICENSE file must specify a MIT, GNU GPL, or all rights reserved license`,
		`Indentation is tab-based and only line feeds (\n) mark the end of lines`,
		`All included files are UTF-8 compatible text files`,
	}

	status := ``
	//run each validator and display output polymorphically
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
			fmt.Println(unit.GetMsg() + "\n")
		}
		if detailMode && unit.GetIssueCt() > 0 {
			fmt.Println(unit.GetIssues() + "\n")
		}
		dispMsg(status, borderLen)
		dispBorder(borderChar, borderLen)
	}
}
