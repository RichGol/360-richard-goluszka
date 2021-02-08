package linefmtchk

import (
	"io/ioutil"
	"os"
	"strings"
)

//LineFmtChecker ... Path string
type LineFmtChecker struct {
	Path, msg string
}

//Validate ... implements validator interface in val.go
func (lfc *LineFmtChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(lfc.Path)
	if err != nil {
		lfc.msg = `Failed to open directory`
		return false
	}

	//step through each entry in lfc.Path directory
	for _, fi := range files {
		if fi.IsDir() { //validate subdirectories
			tmpStr := lfc.Path
			lfc.Path += string(os.PathSeparator) + fi.Name()
			lfc.msg += `Checking: ` + lfc.Path + "\n"
			if lfc.Validate() {
				lfc.Path = tmpStr
				continue
			}
			return false
		}
		//validate files
		filePath := lfc.Path + string(os.PathSeparator) + fi.Name()
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			lfc.msg += `Failed to open file: ` + filePath
			return false
		}
		lfc.msg += `Checking: ` + filePath + "\n"
		//check for invalid line feed character and spaces in-place of tabs
		if strings.Contains(string(content), "\r") || strings.Contains(string(content),
			"\u0020\u0020\u0020\u0020") {
			return false //file fail
		}
	}
	return true //all files pass
}

//GetMsg ... implements validator interface in val.go
func (lfc *LineFmtChecker) GetMsg() string {
	return strings.TrimSuffix(lfc.msg, "\n")
}
