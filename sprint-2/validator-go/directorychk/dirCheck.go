package directorychk

import (
	"io/ioutil"
	"os"
	"strings"
)

//DirChecker ... path, msg string
type DirChecker struct {
	Path, msg string
}

//Validate ... implements validator interface in val.go
func (dc *DirChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(dc.Path)
	if err != nil {
		dc.msg = `Failed to open directory or retrieve file`
		return false
	}

	//validate each file for compliance
	for _, fi := range files {
		tmpName := strings.ToUpper(fi.Name())
		if fi.IsDir() { //validate subdirectories
			tmpStr := dc.Path
			dc.Path = dc.Path + string(os.PathSeparator) + fi.Name()
			dc.msg = dc.msg + `Checking: ` + dc.Path + "\n"
			if dc.Validate() {
				dc.Path = tmpStr
				continue
			}
			return false
		} else if strings.Contains("LICENSE README.MD", tmpName) ||
			strings.Contains(tmpName, ".GO") { //file complies
			dc.msg = dc.msg + `Checking: ` + fi.Name() + "\n"
			continue
		}
		//file fails
		dc.msg = `Directory contains an invalid file: ` + fi.Name()
		return false
	}
	//project complies
	return true
}

//GetMsg ... implements validator interface in val.go
func (dc *DirChecker) GetMsg() string {
	return dc.msg
}
