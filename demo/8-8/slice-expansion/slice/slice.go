package slice

import "unsafe"

type slice struct {
	array unsafe.Pointer //指向底层数组的指针:指向切片引用的底层数组的起始位置
	len   int            //切片的长度：表示切片中元素的个数
	cap   int            //切片的容量：表示底层数组中可访问的元素个数，从切片的第一个元素开始算起。(cap>=len)
}

func growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type) slice {
	// oldLen 为旧的切片底层数组的长度
	oldLen := newLen - num
	if raceenabled {
		callerpc := getcallerpc()
		racereadrangepc(oldPtr, uintptr(oldLen*int(et.size)), callerpc, abi.FuncPCABIInternal(growslice))
	}
	if msanenabled {
		msanread(oldPtr, uintptr(oldLen*int(et.size)))
	}
	if asanenabled {
		asanread(oldPtr, uintptr(oldLen*int(et.size)))
	}

	// 分配的新的长度不能小于 0（整数溢出的时候会是负数）
	if newLen < 0 {
		panic(errorString("growslice: len out of range"))
	}

	// 如果结构或数组类型不包含大小大于零的字段（或元素），则其大小为零。
	//（空数组、空结构体，type b [0]int、type zero struct{}）
	// 两个不同的零大小变量在内存中可能具有相同的地址

	if et.size == 0 {
		// append 不应创建具有 nil 指针但长度非零的切片。
		// 在这种情况下，我们假设 append 不需要保留 oldPtr。
		return slice{unsafe.Pointer(&zerobase), newLen, newLen}
	}

	// newcap 是新切片底层数组的容量
	newcap := oldCap
	// 两倍容量
	doublecap := newcap + newcap
	if newLen > doublecap {
		// 如果追加元素之后，新的切片长度比旧切片 2 倍容量还大，
		// 则将新的切片的容量设置为跟长度一样
		newcap = newLen
	} else {
		const threshold = 256
		if oldCap < threshold {
			// 旧的切片容量小于 256 的时候，进行两倍扩容
			newcap = doublecap
		} else {
			// oldCap >= 256 检查 0<newcap 以检测溢出并防止无限循环。
			for 0 < newcap && newcap < newLen {
				// 从小切片的增长 2 倍过渡到大切片的增长 1.25 倍
				newcap += (newcap + 3*threshold) / 4
			}
			// 当 newcap 计算溢出时，将 newcap 设置为请求的上限
			if newcap <= 0 {
				newcap = newLen
			}
		}
	}

	// 计算实际所需要的内存大小
	// 是否溢出
	var overflow bool
	// lenmem 表示旧的切片长度所需要的内存大小
	//（lenmem 就是将旧切片数据复制到新切片的时候指定需要复制的内存大小）
	// newlenmem 表示新的切片长度所需要的内存大小
	// capmem 表示新的切片容量所需要的内存大小
	var lenmem, newlenmem, capmem uintptr
	//根据 et.size 做一些计算上的优化：
	// 对于 et.size= 1，我们不需要任何除法/乘法。
	// 对于 goarch.PtrSize，编译器会将除法/乘法优化为移位一个常数。
	// 对于 2 的幂，使用可变移位。
	switch {
	case et.size == 1: //// 比如 []byte，所需内存大小 = size
		lenmem = uintptr(oldLen)
		newlenmem = uintptr(newLen)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == goarch.PtrSize:
		lenmem = uintptr(oldLen) * goarch.PtrSize
		newlenmem = uintptr(newLen) * goarch.PtrSize
		capmem = roundupsize(uintptr(newcap) * goarch.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/goarch.PtrSize
		newcap = int(capmem / goarch.PtrSize)
	case isPowerOfTwo(et.size): // 比如 []int64，所需内存大小 = size << shift，也就是 size * 2^shift（2^shift 是 et.size）
		var shift uintptr
		if goarch.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.TrailingZeros64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.TrailingZeros32(uint32(et.size))) & 31
		}
		lenmem = uintptr(oldLen) << shift
		newlenmem = uintptr(newLen) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
		capmem = uintptr(newcap) << shift
	default:
		lenmem = uintptr(oldLen) * et.size
		newlenmem = uintptr(newLen) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
		capmem = uintptr(newcap) * et.size
	}

	// The check of overflow in addition to capmem > maxAlloc is needed
	// to prevent an overflow which can be used to trigger a segfault
	// on 32bit architectures with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	//   s = append(s, d, d, d, d)
	//   print(len(s), "\n")
	// }
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: len out of range"))
	}

	var p unsafe.Pointer
	if et.ptrdata == 0 {
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from oldLen to newLen.
		// Only clear the part that will not be overwritten.
		// The reflect_growslice() that calls growslice will manually clear
		// the region not cleared here.
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			// Only shade the pointers in oldPtr since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(oldPtr), lenmem-et.size+et.ptrdata)
		}
	}
	// 旧切片数据复制到新切片中，复制的内容大小为 lenmem
	//（从 oldPtr 复制到 p）
	memmove(p, oldPtr, lenmem)

	return slice{p, newLen, newcap}
}
