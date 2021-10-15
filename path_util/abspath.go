package path_util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var ExEPath string

func ExEPathPrintln() {
	fmt.Println("EXE abs path:", ExEPath)
}

//return absolute path given relative path to the executable file folder
func GetAbsPath(relpath string) string {
	return filepath.Join(ExEPath, relpath)
}

func init() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err.Error())
	}
	runPath, err := filepath.Abs(file)
	if err != nil {
		panic(err.Error())
	}
	index := strings.LastIndex(runPath, string(os.PathSeparator))
	ExEPath = runPath[:index]
}
