package licensechk

import (
	"io/ioutil"
	"os"
	"strings"
)

//LicenseChecker ... Path string
type LicenseChecker struct {
	Path, msg, filePath string
}

//Validate ... implements validator interface in val.go
func (lc *LicenseChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(lc.Path)
	if err != nil {
		lc.msg = `Failed to open directory or retrieve file`
		return false
	}

	//find LICENSE file in directory
	for _, fi := range files {
		if strings.EqualFold(fi.Name(), "LICENSE") {
			lc.filePath = lc.Path + string(os.PathSeparator) + fi.Name()
			break
		}
		return false
	}

	//validate LICENSE file contents
	bytes, err := ioutil.ReadFile(lc.filePath)
	str := string(bytes)
	if strings.Contains(str, `GNU`) || strings.Contains(str, `MIT`) ||
		strings.Contains(strings.ToUpper(str), `ALL RIGHTS RESERVED`) {
		lc.msg = `Checking: ` + lc.filePath
		return true //file pass
	}
	return false //file fail
}

//GetMsg ... implements validator interface in val.go
func (lc *LicenseChecker) GetMsg() string {
	return strings.TrimSuffix(lc.msg, "\n")
}
