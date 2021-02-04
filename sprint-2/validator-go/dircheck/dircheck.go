package dircheck

import (
	"io/ioutil"
	"os"
	"strings"
)

//DirChecker ... path, msg string
type DirChecker struct {
	Path, Msg string
}

type validator interface {
	validate() bool
}

//Validate ... implementation of Validator interface by DirChecker
func (dc *DirChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(dc.Path)
	if err != nil {
		dc.Msg = `Failed to open directory or retrieve file`
		return false
	}

	//validate each file for compliance
	for _, fi := range files {
		tmpName := strings.ToUpper(fi.Name())
		if fi.IsDir() {
			//validate subdirectories
			dc.Path = dc.Path + string(os.PathSeparator) + fi.Name()
			if dc.Validate() {
				continue
			}
			return false
		} else if strings.Contains("LICENSE README.MD", tmpName) ||
			strings.Contains(tmpName, ".GO") {
			//file complies
			continue
		}
		//file fails
		dc.Msg = `Directory contains an invalid file: ` + fi.Name()
		return false
	}
	//project complies
	return true
}
