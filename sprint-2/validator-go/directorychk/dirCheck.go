package directorychk

import (
	"io/ioutil"
	"os"
	"strings"
)

//DirChecker ... Path string
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

	//step through each entry in dc.Path directory
	for _, fi := range files {
		tmpName := strings.ToUpper(fi.Name())
		if fi.IsDir() { //validate subdirectories
			tmpStr := dc.Path
			dc.Path += string(os.PathSeparator) + fi.Name()
			dc.msg += `Checking: ` + dc.Path + "\n"
			if dc.Validate() {
				dc.Path = tmpStr
				continue
			}
			return false
		} else if strings.Contains("LICENSE README.MD", tmpName) ||
			strings.Contains(tmpName, ".GO") || strings.Contains(tmpName, ".MOD") {
			dc.msg += `Checking: ` + dc.Path + string(os.PathSeparator) + fi.Name() + "\n"
			continue //file pass
		}
		dc.msg += `Directory contains an invalid file: ` + dc.Path + string(os.PathSeparator) +
			fi.Name()
		return false //file fail
	}
	return true //all files pass
}

//GetMsg ... implements validator interface in val.go
func (dc *DirChecker) GetMsg() string {
	return strings.TrimSuffix(dc.msg, "\n")
}
