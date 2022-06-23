package data

import (
	"time"
)

type cData struct {
	localData []byte
	updated   time.Time
}

func checkCache(cache map[string]cData, stockSymbol string) bool {
	if c, ok := cache[stockSymbol]; ok {
		if time.Now().Sub(c.updated) < (time.Hour * 24) {
			return true
		}
	}
	return false
}
