package utf8chk

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

//UTF8Checker .. Path string
type UTF8Checker struct {
	Path    string
	msg     string
	issues  string
	issueCt int
}

//Validate ... implements validator interface in val.go
func (uc *UTF8Checker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(uc.Path)
	if err != nil {
		uc.msg += `Failed to open directory` + uc.Path + "\n"
		return false
	}

	status := true
	tmpPath := ``
	filePath := ``
	//step through each entry in uc.Path directory
	for _, fi := range files {
		filePath = uc.Path + string(os.PathSeparator) + fi.Name()
		uc.msg += `Checking: ` + filePath + "\n"

		if fi.IsDir() { //validate subdirectories
			tmpPath = uc.Path
			uc.Path += string(os.PathSeparator) + fi.Name()
			if !uc.Validate() {
				status = false
			}
			uc.Path = tmpPath
			continue
		}
		//open file
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			uc.msg += `Failed to open file: ` + filePath + "\n"
			return false
		}

		//Check each line for UTF-8 validity
		for lineNum, line := range strings.Split(string(content), "\n") {
			if !utf8.ValidString(line) {
				uc.issues += fmt.Sprintf("Line %d in File: %s\tnon-utf8 text\n", lineNum, line)
				uc.issueCt++
				status = false //file fail
			}
		}
	}
	return status
}

//GetMsg ... implements validator interface in val.go
func (uc *UTF8Checker) GetMsg() string {
	return strings.TrimSuffix(uc.msg, "\n")
}

//GetIssues ... implements validator interface in val.go
func (uc *UTF8Checker) GetIssues() string {
	return strings.TrimSuffix(uc.issues, "\n")
}

//GetIssueCt ... implements validator interface in val.go
func (uc *UTF8Checker) GetIssueCt() int {
	return uc.issueCt
}
