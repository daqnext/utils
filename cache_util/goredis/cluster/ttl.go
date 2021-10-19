package cluster

import (
	"context"
	"math/rand"

	"github.com/go-redis/redis/v8"
)

//for redis cluster , when return true , you need to refresh the key-value
func CheckTtlRefresh(Redis *redis.ClusterClient, ctx context.Context, keystr string) bool {

	if rand.Intn(10) == 1 {
		Duration, error := Redis.TTL(ctx, keystr).Result()
		if error != nil {
			return true
		}
		secleft := Duration.Seconds()
		if secleft > 0 && secleft < 10 {
			if rand.Intn(int(secleft)*5) == 1 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}
