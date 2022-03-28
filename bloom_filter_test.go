/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2022/3/28 下午4:40
 * @note:
 */

package bloom

import (
	"strconv"
	"testing"
)

func BenchmarkBloomFilter_Add(b *testing.B) {
	b.StopTimer()
	filter := New(1 << 30)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		filter.Add([]byte(strconv.Itoa(i)))
	}
}

func BenchmarkBloomFilter_Test(b *testing.B) {
	b.StopTimer()
	filter := New(1 << 30)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = filter.Test([]byte(strconv.Itoa(i)))
	}
}
