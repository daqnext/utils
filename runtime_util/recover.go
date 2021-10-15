package runtime_util

import "runtime"

func RecoverToString(err interface{}) string {
	var ErrStr string
	switch e := err.(type) {
	case string:
		ErrStr = e
	case runtime.Error:
		ErrStr = e.Error()
	case error:
		ErrStr = e.Error()
	default:
		ErrStr = "recover default err"
	}
	return ErrStr
}
