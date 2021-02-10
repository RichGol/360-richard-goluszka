package licensechk

import (
	"io/ioutil"
	"os"
	"strings"
)

//LicenseChecker ... Path string
type LicenseChecker struct {
	Path, filePath string
	msg, issues    string
	issueCt        int
}

//Validate ... implements validator interface in val.go
func (lc *LicenseChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(lc.Path)
	if err != nil {
		lc.msg += `Failed to open directory: ` + lc.Path + "\n"
		return false
	}

	//find LICENSE file in directory
	for _, fi := range files {
		if strings.EqualFold(fi.Name(), `LICENSE`) {
			lc.filePath = lc.Path + string(os.PathSeparator) + fi.Name()
			break
		}
		lc.issueCt++
		return false
	}

	//open LICENSE file
	bytes, err := ioutil.ReadFile(lc.filePath)
	if err != nil {
		lc.msg += `Failed to open file: ` + lc.filePath + "\n"
		return false
	}

	//validate contents
	lc.msg += `Checking: ` + lc.filePath + "\n"
	str := string(bytes) //TODO: Modify to check by line
	if strings.Contains(str, `GNU`) || strings.Contains(str, `MIT`) ||
		strings.Contains(strings.ToUpper(str), `ALL RIGHTS RESERVED`) {
		return true //file pass
	}
	lc.issueCt++
	lc.issues += `LICENSE File Invalid: ` + lc.filePath + "\n"
	lc.issues += `  --does not contain "GNU", "MIT", or "all rights reserved"` + "\n"
	return false //file fail
}

//GetMsg ... implements validator interface in val.go
func (lc *LicenseChecker) GetMsg() string {
	return strings.TrimSuffix(lc.msg, "\n")
}

//GetIssues ... implements validator interface in val.go
func (lc *LicenseChecker) GetIssues() string {
	return strings.TrimSuffix(lc.issues, "\n")
}

//GetIssueCt ... implements validator interface in val.go
func (lc *LicenseChecker) GetIssueCt() int {
	return lc.issueCt
}
