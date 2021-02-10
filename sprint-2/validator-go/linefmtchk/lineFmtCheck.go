package linefmtchk

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

//LineFmtChecker ... Path string
type LineFmtChecker struct {
	Path    string
	msg     string
	issues  string
	issueCt int
}

//Validate ... implements validator interface in val.go
func (lfc *LineFmtChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(lfc.Path)
	if err != nil {
		lfc.msg += `Failed to open directory: ` + lfc.Path + "\n"
		return false
	}

	//create regular expression to validate tab requirement
	regExpr := regexp.MustCompile(`^(    )+[a-z|A-Z]+`)

	status := true
	tmpPath := ``
	filePath := ``
	//step through each entry in lfc.Path directory
	for _, fi := range files {
		filePath = lfc.Path + string(os.PathSeparator) + fi.Name()
		lfc.msg += `Checking: ` + filePath + "\n"

		if fi.IsDir() { //validate subdirectories
			tmpPath = lfc.Path
			lfc.Path += string(os.PathSeparator) + fi.Name()
			if !lfc.Validate() {
				status = false
			}
			lfc.Path = tmpPath
			continue
		}

		//open file
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			lfc.msg += `Failed to open file: ` + filePath + "\n"
			return false
		}

		//check each line for "\r" and space-indenting
		for lineNum, line := range strings.Split(string(content), "\n") {
			if strings.Contains(line, "\r") {
				lfc.issues += fmt.Sprintf("Issue: %s has wrong line feeds (line %d)\n", filePath,
					lineNum+1)
			} else if regExpr.MatchString(line) {
				lfc.issues += fmt.Sprintf("Issue: %s has space-indentation (line %d)\n", filePath,
					lineNum+1)
			} else {
				continue //file pass
			}
			lfc.issueCt++
			status = false //file fail
		}
	}
	return status
}

//GetMsg ... implements validator interface in val.go
func (lfc *LineFmtChecker) GetMsg() string {
	return strings.TrimSuffix(lfc.msg, "\n")
}

//GetIssues ... implements validator interface in val.go
func (lfc *LineFmtChecker) GetIssues() string {
	return strings.TrimSuffix(lfc.issues, "\n")
}

//GetIssueCt ... implements validator interface in val.go
func (lfc *LineFmtChecker) GetIssueCt() int {
	return lfc.issueCt
}
