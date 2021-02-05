package directorychk

import (
	"io/ioutil"
	"os"
	"strings"
)

//DirChecker ... path, msg string
type DirChecker struct {
	Path, Msg string
}

//Validate ... implements validator interface in val.go
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
		if fi.IsDir() { //validate subdirectories
			tmpStr := dc.Path
			dc.Path = dc.Path + string(os.PathSeparator) + fi.Name()
			dc.Msg = dc.Msg + `Checking: ` + dc.Path + "\n"
			if dc.Validate() {
				dc.Path = tmpStr
				continue
			}
			return false
		} else if strings.Contains("LICENSE README.MD", tmpName) ||
			strings.Contains(tmpName, ".GO") { //file complies
			dc.Msg = dc.Msg + `Checking: ` + fi.Name() + "\n"
			continue
		}
		//file fails
		dc.Msg = `Directory contains an invalid file: ` + fi.Name()
		return false
	}
	//project complies
	return true
}

//GetMsg ... implements validator interface in val.go
func (dc *DirChecker) GetMsg() string {
	return dc.Msg
}
