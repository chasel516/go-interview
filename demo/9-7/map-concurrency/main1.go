package main

import (
	"log"
	"strconv"
	"sync"
)

const ShardCount = 32

// 将map分成SHARD_COUNT个分片，每个分片作为切片的一个元素
type ConcurrentShardMap []*ConcurrentMapShared

// 对每个分片上map进行加锁
type ConcurrentMapShared struct {
	items        map[string]interface{}
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// 创建map
func NewConcurrentShardMap() ConcurrentShardMap {
	m := make(ConcurrentShardMap, ShardCount)
	for i := 0; i < ShardCount; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

// 根据key计算分片索引
func (m ConcurrentShardMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv(key))%uint(ShardCount)]
}

// FNV-1a算法是一种简单且快速的哈希算法，特别适合对字符串进行哈希计算
// 这里的目的是将字符串映射成整数
func fnv(key string) uint64 {
	var h uint64 = 14695981039346656037 // offset
	for i := 0; i < len(key); i++ {
		h = h ^ uint64(key[i])
		h = h * 1099511628211 // prime
	}
	return h
}

func (m ConcurrentShardMap) Set(key string, value interface{}) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.Lock() //对这个分片加锁，执行业务操作
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentShardMap) Get(key string) (interface{}, bool) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.RLock()
	// 从这个分片读取key的值
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// 使用回调函数作为迭代器，以更低的代价获取全部元素
func (m ConcurrentShardMap) Range(f func(key, value any) bool) {
	for index, _ := range m {
		shard := m[index]
		shard.RLock()
		for key, value := range shard.items {
			if !f(key, value) {
				shard.RUnlock()
				return
			}
		}
		shard.RUnlock()
	}
}

func main() {
	m := NewConcurrentShardMap()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			m.Set(strconv.Itoa(index), index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value any) bool {
		log.Println(key, value)
		return true
	})
}
