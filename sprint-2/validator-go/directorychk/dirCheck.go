package directorychk

import (
	"io/ioutil"
	"os"
	"strings"
)

//DirChecker ... Path string
type DirChecker struct {
	Path    string
	msg     string
	issues  string
	issueCt int
}

//Validate ... implements validator interface in val.go
func (dc *DirChecker) Validate() bool {
	//open directory and get files
	files, err := ioutil.ReadDir(dc.Path)
	if err != nil {
		dc.msg += `Failed to open directory: ` + dc.Path + "\n"
		return false
	}

	status := true
	tmpPath := ``
	fileName := ``
	//step through each entry in dc.Path directory
	for _, fi := range files {
		fileName = strings.ToUpper(fi.Name())
		dc.msg += `Checking: ` + dc.Path + string(os.PathSeparator) + fi.Name() + "\n"

		if fi.IsDir() { //validate subdirectories
			tmpPath = dc.Path
			dc.Path += string(os.PathSeparator) + fi.Name()
			if !dc.Validate() {
				status = false
			}
			dc.Path = tmpPath
			continue
		} else if strings.Contains(fileName, `.EXE`) {
			continue
		} else if strings.Contains("LICENSE README.MD", fileName) ||
			strings.Contains(fileName, ".GO") || strings.Contains(fileName, ".MOD") {
			continue //file pass
		} else {
			dc.issues += `Issue: ` + dc.Path + string(os.PathSeparator) + fi.Name() +
				" is non-project file\n"
			dc.issueCt++ //file fail
			status = false
		}
	}
	return status
}

//GetMsg ... implements validator interface in val.go
func (dc *DirChecker) GetMsg() string {
	return strings.TrimSuffix(dc.msg, "\n")
}

//GetIssues ... implements validator interface in val.go
func (dc *DirChecker) GetIssues() string {
	return strings.TrimSuffix(dc.issues, "\n")
}

//GetIssueCt ... implements validator interface in val.go
func (dc *DirChecker) GetIssueCt() int {
	return dc.issueCt
}
