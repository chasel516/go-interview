package main

import "unsafe"

const (
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

// 定义了map的结构
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // 元素的个数, len() 的值
	flags     uint8
	B         uint8  //bucket个数为：2^B；可以保存的元素个数：填充因子(默认6.5) * 2^B items)
	noverflow uint16 /// 溢出桶数量
	hash0     uint32 // 哈希因子

	buckets    unsafe.Pointer // Buckets数组，大小为 2^B
	oldbuckets unsafe.Pointer // 发生扩容前的Buckets，在增长时非nil
	nevacuate  uintptr        // 迁移状态，进度

	extra *mapextra // 用于存储一些不是所有哈希表都需要的字段
}

type mapextra struct {
	//指向bmap类型切片的指针，bmap类型是哈希表中存储数据的基本单元。
	//overflow字段用于存储所有溢出桶的指针，溢出桶是当哈希表中某个位置发生冲突时，用于存放多余数据的额外桶。
	//overflow字段只有在哈希表的键和值都不包含指针并且可以内联时才使用，这样可以避免扫描这些哈希表。
	overflow *[]*bmap
	//这也是一个指向bmap类型切片的指针，它用于存储旧哈希表中的溢出桶的指针，
	//旧哈希表是当哈希表需要扩容时，原来的哈希表称为旧哈希表，新分配的哈希表称为新哈希表。
	//oldoverflow字段也只有在键和值都不包含指针并且可以内联时才使用。
	oldoverflow *[]*bmap

	// 指向bmap类型的指针，它用于存储一个空闲的溢出桶，当需要分配一个新的溢出桶时，就从nextOverflow中取出一个，
	//如果nextOverflow为空，则从堆上分配一个新的溢出桶。
	nextOverflow *bmap
}

// 定义了hmap.buckets中每个bucket的结构
type bmap struct {
	//bucketCnt 是常量=8，一个bucket最多存储8个key/value对
	tophash [bucketCnt]uint8
}
