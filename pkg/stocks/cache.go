package stocks

import (
	"time"
)

type cacheData struct {
	data        []byte
	lastUpdated time.Time
}

func checkCache(cache map[string]cacheData, stockSymbol string) bool {
	now := time.Now()
	if c, ok := cache[stockSymbol]; ok {
		if now.Sub(c.lastUpdated) < (time.Hour * 24) {
			return true
		}
	}
	return false
}
