package licensechk

import (
	"io/ioutil"
	"os"
	"strings"
)

//LicenseChecker ... struct
type LicenseChecker struct {
	Path, Msg, fileLoc string
}

//Validate ... implements validator interface in val.go
func (lc *LicenseChecker) Validate() bool {
	//find the license file
	files, err := ioutil.ReadDir(lc.Path)
	if err != nil {
		lc.Msg = `Failed to open directory or retrieve file`
		return false
	}

	for _, fi := range files {
		if strings.EqualFold(fi.Name(), "LICENSE") {
			lc.fileLoc = lc.Path + string(os.PathSeparator) + fi.Name()
			break
		}
		return false
	}

	bytes, err := ioutil.ReadFile(lc.fileLoc)
	str := string(bytes)
	if strings.Contains(str, `GNU`) || strings.Contains(str, `MIT`) ||
		strings.Contains(strings.ToUpper(str), `ALL RIGHTS RESERVED`) {
		lc.Msg = `Checking ` + lc.fileLoc
		return true
	}
	return false
}

//GetMsg ... implements validator interface in val.go
func (lc *LicenseChecker) GetMsg() string {
	return lc.Msg
}
