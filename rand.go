package mathgen

import (
	"math/rand"
	"time"
)

// r is default random instance
var r *rand.Rand

// getRand will get default random instance. Create new one in case it's unvailable
func getRand() *rand.Rand {
	if r == nil {
		t := time.Now().UnixNano()
		r = rand.New(rand.NewSource(t))
	}
	return r
}
