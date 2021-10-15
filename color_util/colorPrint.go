package color_util

import (
	"fmt"
	"runtime"
)

const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Purple = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"

func ColorPrintln(color string, a ...interface{}) (n int, err error) {
	if runtime.GOOS == "windows" {
		return fmt.Println(a...)
	} else {
		a = append([]interface{}{color}, a...)
		return fmt.Println(a...)
	}
}

func ColorPrint(color string, a ...interface{}) (n int, err error) {
	if runtime.GOOS == "windows" {
		return fmt.Print(a...)
	} else {
		a = append([]interface{}{color}, a...)
		return fmt.Print(a...)
	}
}

func ColorPrintf(color string, format string, a ...interface{}) (n int, err error) {
	if runtime.GOOS == "windows" {
		return fmt.Printf(format, a...)
	} else {
		return fmt.Printf(color+format, a...)
	}
}
