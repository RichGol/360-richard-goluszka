package licensechk

import (
	"io/ioutil"
	"os"
	"strings"
)

//LicenseChecker ... Path string
type LicenseChecker struct {
	Path        string
	filePath    string
	msg, issues string
	issueCt     int
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
		lc.issues += `Issue: No LICENSE file`
		return false
	}

	//open LICENSE file
	bytes, err := ioutil.ReadFile(lc.filePath)
	if err != nil {
		lc.msg += `Failed to open file: ` + lc.filePath + "\n"
		return false
	}
	lc.msg += `Checking: ` + lc.filePath + "\n"

	//validate contents
	content := strings.ToUpper(string(bytes))
	licMIT := `PERMISSION IS HEREBY GRANTED, FREE OF CHARGE,`
	licReserved := `ALL RIGHTS RESERVED`
	if strings.Contains(content, `GNU`) || strings.Contains(content, `MIT`) ||
		strings.Contains(content, licMIT) || strings.Contains(content, licReserved) {
		return true //file pass
	}
	lc.issueCt++
	lc.issues += `Issue: ` + lc.filePath +
		" does not mention GNU, MIT, or all rights reserved\n"
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
