package cache_util

import "math/rand"

func CheckTtlRefresh(secleft int64) bool {
	if secleft > 0 && secleft < 4 {
		if rand.Intn(100) == 50 {
			return true
		}
	}
	return false
}
