package popcount

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// TODO(allen):时间测试有问题
func TestPopCount(t *testing.T) {
	ts := time.Now()
	rand.Seed(time.Now().Unix())

	xs := []uint64{646782549, 234234, 759473, 573845, 123, 8594, 45789435, 549, 548, 12380, 345894, 56379, 1238, 67, 213, 8098235}
	for j := 1; j < 10000; j++ {
		xs = append(xs, rand.Uint64())
	}

	ts = time.Now()
	for x := range xs {
		fmt.Print(PopCount23(uint64(x)))
	}
	fmt.Println()
	fmt.Println("PopCount23 spend time", time.Since(ts).Nanoseconds())

	ts = time.Now()
	for x := range xs {
		fmt.Print(PopCount24(uint64(x)))
	}
	fmt.Println()
	fmt.Println("PopCount24 spend time", time.Since(ts).Nanoseconds())

	ts = time.Now()
	for x := range xs {
		fmt.Print(PopCount25(uint64(x)))
	}
	fmt.Println()
	fmt.Println("PopCount25 spend time", time.Since(ts).Nanoseconds())

	ts = time.Now()
	for x := range xs {
		fmt.Print(PopCount(uint64(x)))
	}
	fmt.Println()
	fmt.Println("PopCount spend time", time.Since(ts).Nanoseconds())

}
