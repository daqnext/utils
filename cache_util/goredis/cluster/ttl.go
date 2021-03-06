package cluster

import (
	"context"
	"math/rand"
	"time"

	gofastcache "github.com/daqnext/go-fast-cache"
	"github.com/daqnext/utils/cache_util"
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

func SmartCheck_LocalCache_Redis(ctx context.Context, Redis *redis.ClusterClient, LocalCache *gofastcache.LocalCache, keystr string) (interface{}, int64, bool) {
	localvalue, ttl, localexist := LocalCache.Get(keystr)
	if !cache_util.CheckTtlRefresh(ttl) && localexist {
		randSyncStr := keystr + ":randsync"
		rresult, err := Redis.Get(ctx, randSyncStr).Result()
		if err == nil && rresult == LocalCache.GetRand(randSyncStr) {
			return localvalue, ttl, true
		}
	}
	return nil, 0, false
}

func SmartSet_LocalCache_Redis(ctx context.Context, Redis *redis.ClusterClient, LocalCache *gofastcache.LocalCache, keystr string, value interface{}, ttlSecond int64) {
	LocalCache.Set(keystr, value, ttlSecond)
	randSyncStr := keystr + ":randsync"
	strsrc := LocalCache.SetRand(randSyncStr, ttlSecond+10)
	Redis.Set(ctx, randSyncStr, strsrc, time.Duration(ttlSecond+30)*time.Second)
}

func SmartDel_LocalCache_Redis(ctx context.Context, Redis *redis.ClusterClient, LocalCache *gofastcache.LocalCache, keystr string) {
	randSyncStr := keystr + ":randsync"
	LocalCache.Delete(keystr)
	LocalCache.Delete(randSyncStr)
	Redis.Del(ctx, keystr)
	Redis.Del(ctx, randSyncStr)
}
