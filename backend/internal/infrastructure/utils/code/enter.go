package code

import (
	"math/rand"
	"sync"
	"time"
)

// 每个 goroutine 自己拿一份私有随机源，避免锁
var pool = sync.Pool{
	New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

func GenCode() string {
	r := pool.Get().(*rand.Rand)
	defer pool.Put(r)

	var nums [6]byte
	for i := 0; i < 6; i++ {
		nums[i] = byte(r.Intn(10)) + '0' // 0-9 -> '0'-'9'
	}
	return string(nums[:])
}
