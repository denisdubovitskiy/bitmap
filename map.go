package bitmap

import (
	"math"
	"sync"
)

const bucketSize = 8 // to lock small pieces

type Bitmap interface {
	Set(n int)
	Clear(n int)
	Has(n int) bool
}

func New(size int) Bitmap {
	bucketsCount := int(math.Ceil(float64(size) / bucketSize))
	if bucketsCount == 0 {
		bucketsCount = 1
	}

	return &bitmap{
		buckets: make([]uint8, bucketsCount),
		muxes:   make([]sync.RWMutex, bucketsCount),
	}
}

type bitmap struct {
	muxes   []sync.RWMutex
	buckets []uint8
}

func (c bitmap) Set(n int) {
	b := int(float64(n) / bucketSize)
	i := uint(math.Abs(float64((b * bucketSize) - n)))

	c.muxes[b].Lock()
	defer c.muxes[b].Unlock()

	c.buckets[b] |= 1 << i
}

func (c bitmap) Clear(n int) {
	b := int(float64(n) / bucketSize)
	i := uint(math.Abs(float64((b * bucketSize) - n)))

	c.muxes[b].Lock()
	defer c.muxes[b].Unlock()

	c.buckets[b] &= ^(1 << i)
}

func (c *bitmap) Has(n int) bool {
	b := int(float64(n) / bucketSize)
	i := uint(math.Abs(float64((b * bucketSize) - n)))

	c.muxes[b].RLock()
	defer c.muxes[b].RUnlock()

	return c.buckets[b]&(1<<i) > 0
}
