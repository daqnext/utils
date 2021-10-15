package runtime_util

import (
	"runtime/debug"
	"strings"
)

func StackString(linelimit int) []string {
	var stacks []string
	if linelimit <= 0 {
		return stacks
	}
	lines := strings.Split(string(debug.Stack()), "\n")
	counter := 0
	for i := 2; i < len(lines); i = i + 2 {
		if counter >= linelimit {
			return stacks
		}
		fomatstr := strings.ReplaceAll(lines[i], "	", "")
		stacks = append(stacks, fomatstr)
		counter++
	}
	return stacks
}
