package utf8chk

import (
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

//UTF8Checker .. Path string
type UTF8Checker struct {
	Path, msg string
}

//Validate ... implements validator interface in val.go
func (uc *UTF8Checker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(uc.Path)
	if err != nil {
		uc.msg = `Failed to open directory`
		return false
	}

	//step through each entry in uc.Path directory
	for _, fi := range files {
		if fi.IsDir() { //validate subdirectories
			tmpStr := uc.Path
			uc.Path += string(os.PathSeparator) + fi.Name()
			uc.msg += `Checking: ` + uc.Path + "\n"
			if uc.Validate() {
				uc.Path = tmpStr
				continue
			}
			return false
		}
		//validate files
		filePath := uc.Path + string(os.PathSeparator) + fi.Name()
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			uc.msg += `Failed to open file: ` + filePath
			return false
		}
		uc.msg += `Checking: ` + filePath + "\n"
		//Check for utf8 validity
		if utf8.Valid(content) {
			continue
		}
		return false
	}
	return true
}

//GetMsg ... implements validator interface in val.go
func (uc *UTF8Checker) GetMsg() string {
	return strings.TrimSuffix(uc.msg, "\n")
}
