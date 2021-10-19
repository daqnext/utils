package cache_util

import "math/rand"

//used for localCache
func CheckTtlRefresh(secleft int64) bool {
	if secleft > 0 && secleft < 8 {
		if rand.Intn(int(secleft)*50) == 1 {
			return true
		}
	}
	return false
}
