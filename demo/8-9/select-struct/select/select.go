package runtim

import (
	"unsafe"
)

// 源码位置：$GOROOT/src/runtime/select.go
const debugSelect = false

// 所有的case语句一起组成scase的结构数组
type scase struct {
	c    *hchan         // chan的结构体类型指针，如果是 default 分支则为 nil
	elem unsafe.Pointer // //读或者写的缓冲区地址
}

// 源码位置：$GOPATH/src/runtime/chan.go
type hchan struct {
	//当前队列中元素的个数。当我们向channel发送数据时，qcount会增加1；当我们从channel接收数据时，qcount会减少1
	qcount uint

	//如果我们在创建channel时指定了缓冲区的大小，那么dataqsiz就等于指定的大小；否则，dataqsiz为0，表示该channel没有缓冲区。
	dataqsiz uint

	//buf字段是一个unsafe.Pointer类型的指针，指向缓冲区的起始地址。如果该channel没有缓冲区，则buf为nil。
	buf unsafe.Pointer

	//表示缓冲区中每个元素的大小。当我们创建channel时，Golang会根据元素的类型计算出elemsize的值。
	elemsize uint16

	// channel 是否已经关闭,当我们通过close函数关闭一个channel时，Golang会将closed字段设置为true。
	closed uint32

	//表示下一次接收元素的位置.当我们从channel接收数据时，Golang会从缓冲区中recvx索引的位置读取数据，并将recvx加1
	recvx uint

	//表示下一次发送元素的位置。在channel的发送操作中，如果缓冲区未满，则会将数据写入到sendx指向的位置，并将sendx加1。
	//如果缓冲区已满，则发送操作会被阻塞，直到有足够的空间可用。
	sendx uint

	// 等待接收数据的 goroutine 队列，用于存储等待从channel中读取数据的goroutine。
	//当channel中没有数据可读时，接收者goroutine会进入recvq等待队列中等待数据的到来。
	//当发送者goroutine写入数据后，会将recvq等待队列中的接收者goroutine唤醒，并进行读取操作。
	//在进行读取操作时，会先检查recvq等待队列是否为空，如果不为空，则会将队列中的第一个goroutine唤醒进行读取操作。
	//同时，由于recvq等待队列是一个FIFO队列，因此等待时间最长的goroutine会排在队列的最前面，最先被唤醒进行读取操作。
	recvq waitq

	// 等待发送数据的 goroutine 队列。sendq 字段是一个指向 waitq 结构体的指针，waitq 是一个用于等待队列的结构体。
	//waitq 中包含了一个指向等待队列中第一个协程的指针和一个指向等待队列中最后一个协程的指针。
	//当一个协程向一个 channel 中发送数据时，如果该 channel 中没有足够的缓冲区来存储数据，那么发送操作将会被阻塞，
	//直到有另一个协程来接收数据或者 channel 中有足够的缓冲区来存储数据。当一个协程被阻塞在发送操作时，
	//它将会被加入到 sendq 队列中，等待另一个协程来接收数据或者 channel 中有足够的缓冲区来存储数据。
	sendq waitq

	//channel的读写锁，确保多个gorutine同时访问时的并发安全，保证读写操作的原子性和互斥性。
	//当一个goroutine想要对channel进行读写操作时，首先需要获取lock锁。如果当前lock锁已经被其他goroutine占用，
	//则该goroutine会被阻塞，直到lock锁被释放。一旦该goroutine获取到lock锁，就可以进行读写操作，并且在操作完成后释放lock锁，
	//以便其他goroutine可以访问channel底层数据结构。
	lock mutex
}

// selectgo函数实现了Go语言中的select语句。
//遍历所有的case语句，如果所有的case都未就绪，则走default，如果没有default，则会阻塞。
//如果有就绪channel，则直接跳出循环进行管道操作并返回

// 第一个返回值返回的是被选择的scase的索引，这个索引就是需要执行的case分支；
// 第一个返回值标记这个case语句是否接收到了一个值。
func selectgo(cas0 *scase, order0 *uint16, pc0 *uintptr, nsends, nrecvs int, block bool) (int, bool) {
	if debugSelect {
		print("select: cas0=", cas0, "\n")
	}

	// NOTE: In order to maintain a lean stack size, the number of scases
	// is capped at 65536.
	cas1 := (*[1 << 16]scase)(unsafe.Pointer(cas0))
	order1 := (*[1 << 17]uint16)(unsafe.Pointer(order0))

	ncases := nsends + nrecvs
	scases := cas1[:ncases:ncases]
	pollorder := order1[:ncases:ncases]
	lockorder := order1[ncases:][:ncases:ncases]
	// NOTE: pollorder/lockorder's underlying array was not zero-initialized by compiler.

	// Even when raceenabled is true, there might be select
	// statements in packages compiled without -race (e.g.,
	// ensureSigM in runtime/signal_unix.go).
	var pcs []uintptr
	if raceenabled && pc0 != nil {
		pc1 := (*[1 << 16]uintptr)(unsafe.Pointer(pc0))
		pcs = pc1[:ncases:ncases]
	}
	casePC := func(casi int) uintptr {
		if pcs == nil {
			return 0
		}
		return pcs[casi]
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	// The compiler rewrites selects that statically have
	// only 0 or 1 cases plus default into simpler constructs.
	// The only way we can end up with such small sel.ncase
	// values here is for a larger select in which most channels
	// have been nilled out. The general code handles those
	// cases correctly, and they are rare enough not to bother
	// optimizing (and needing to test).

	// generate permuted order
	norder := 0
	for i := range scases {
		cas := &scases[i]

		// Omit cases without channels from the poll and lock orders.
		if cas.c == nil {
			cas.elem = nil // allow GC
			continue
		}

		j := fastrandn(uint32(norder + 1))
		pollorder[norder] = pollorder[j]
		pollorder[j] = uint16(i)
		norder++
	}
	pollorder = pollorder[:norder]
	lockorder = lockorder[:norder]

	// sort the cases by Hchan address to get the locking order.
	// simple heap sort, to guarantee n log n time and constant stack footprint.
	for i := range lockorder {
		j := i
		// Start with the pollorder to permute cases on the same channel.
		c := scases[pollorder[i]].c
		for j > 0 && scases[lockorder[(j-1)/2]].c.sortkey() < c.sortkey() {
			k := (j - 1) / 2
			lockorder[j] = lockorder[k]
			j = k
		}
		lockorder[j] = pollorder[i]
	}
	for i := len(lockorder) - 1; i >= 0; i-- {
		o := lockorder[i]
		c := scases[o].c
		lockorder[i] = lockorder[0]
		j := 0
		for {
			k := j*2 + 1
			if k >= i {
				break
			}
			if k+1 < i && scases[lockorder[k]].c.sortkey() < scases[lockorder[k+1]].c.sortkey() {
				k++
			}
			if c.sortkey() < scases[lockorder[k]].c.sortkey() {
				lockorder[j] = lockorder[k]
				j = k
				continue
			}
			break
		}
		lockorder[j] = o
	}

	if debugSelect {
		for i := 0; i+1 < len(lockorder); i++ {
			if scases[lockorder[i]].c.sortkey() > scases[lockorder[i+1]].c.sortkey() {
				print("i=", i, " x=", lockorder[i], " y=", lockorder[i+1], "\n")
				throw("select: broken sort")
			}
		}
	}

	// lock all the channels involved in the select
	sellock(scases, lockorder)

	var (
		gp     *g
		sg     *sudog
		c      *hchan
		k      *scase
		sglist *sudog
		sgnext *sudog
		qp     unsafe.Pointer
		nextp  **sudog
	)

	// pass 1 - look for something already waiting
	var casi int
	var cas *scase
	var caseSuccess bool
	var caseReleaseTime int64 = -1
	var recvOK bool
	for _, casei := range pollorder {
		casi = int(casei)
		cas = &scases[casi]
		c = cas.c

		if casi >= nsends {
			sg = c.sendq.dequeue()
			if sg != nil {
				goto recv
			}
			if c.qcount > 0 {
				goto bufrecv
			}
			if c.closed != 0 {
				goto rclose
			}
		} else {
			if raceenabled {
				racereadpc(c.raceaddr(), casePC(casi), chansendpc)
			}
			if c.closed != 0 {
				goto sclose
			}
			sg = c.recvq.dequeue()
			if sg != nil {
				goto send
			}
			if c.qcount < c.dataqsiz {
				goto bufsend
			}
		}
	}

	if !block {
		selunlock(scases, lockorder)
		casi = -1
		goto retc
	}

	// pass 2 - enqueue on all chans
	gp = getg()
	if gp.waiting != nil {
		throw("gp.waiting != nil")
	}
	nextp = &gp.waiting
	for _, casei := range lockorder {
		casi = int(casei)
		cas = &scases[casi]
		c = cas.c
		sg := acquireSudog()
		sg.g = gp
		sg.isSelect = true
		// No stack splits between assigning elem and enqueuing
		// sg on gp.waiting where copystack can find it.
		sg.elem = cas.elem
		sg.releasetime = 0
		if t0 != 0 {
			sg.releasetime = -1
		}
		sg.c = c
		// Construct waiting list in lock order.
		*nextp = sg
		nextp = &sg.waitlink

		if casi < nsends {
			c.sendq.enqueue(sg)
		} else {
			c.recvq.enqueue(sg)
		}
	}

	// wait for someone to wake us up
	gp.param = nil
	// Signal to anyone trying to shrink our stack that we're about
	// to park on a channel. The window between when this G's status
	// changes and when we set gp.activeStackChans is not safe for
	// stack shrinking.
	gp.parkingOnChan.Store(true)
	gopark(selparkcommit, nil, waitReasonSelect, traceBlockSelect, 1)
	gp.activeStackChans = false

	sellock(scases, lockorder)

	gp.selectDone.Store(0)
	sg = (*sudog)(gp.param)
	gp.param = nil

	// pass 3 - dequeue from unsuccessful chans
	// otherwise they stack up on quiet channels
	// record the successful case, if any.
	// We singly-linked up the SudoGs in lock order.
	casi = -1
	cas = nil
	caseSuccess = false
	sglist = gp.waiting
	// Clear all elem before unlinking from gp.waiting.
	for sg1 := gp.waiting; sg1 != nil; sg1 = sg1.waitlink {
		sg1.isSelect = false
		sg1.elem = nil
		sg1.c = nil
	}
	gp.waiting = nil

	for _, casei := range lockorder {
		k = &scases[casei]
		if sg == sglist {
			// sg has already been dequeued by the G that woke us up.
			casi = int(casei)
			cas = k
			caseSuccess = sglist.success
			if sglist.releasetime > 0 {
				caseReleaseTime = sglist.releasetime
			}
		} else {
			c = k.c
			if int(casei) < nsends {
				c.sendq.dequeueSudoG(sglist)
			} else {
				c.recvq.dequeueSudoG(sglist)
			}
		}
		sgnext = sglist.waitlink
		sglist.waitlink = nil
		releaseSudog(sglist)
		sglist = sgnext
	}

	if cas == nil {
		throw("selectgo: bad wakeup")
	}

	c = cas.c

	if debugSelect {
		print("wait-return: cas0=", cas0, " c=", c, " cas=", cas, " send=", casi < nsends, "\n")
	}

	if casi < nsends {
		if !caseSuccess {
			goto sclose
		}
	} else {
		recvOK = caseSuccess
	}

	if raceenabled {
		if casi < nsends {
			raceReadObjectPC(c.elemtype, cas.elem, casePC(casi), chansendpc)
		} else if cas.elem != nil {
			raceWriteObjectPC(c.elemtype, cas.elem, casePC(casi), chanrecvpc)
		}
	}
	if msanenabled {
		if casi < nsends {
			msanread(cas.elem, c.elemtype.Size_)
		} else if cas.elem != nil {
			msanwrite(cas.elem, c.elemtype.Size_)
		}
	}
	if asanenabled {
		if casi < nsends {
			asanread(cas.elem, c.elemtype.Size_)
		} else if cas.elem != nil {
			asanwrite(cas.elem, c.elemtype.Size_)
		}
	}

	selunlock(scases, lockorder)
	goto retc

bufrecv:
	// can receive from buffer
	if raceenabled {
		if cas.elem != nil {
			raceWriteObjectPC(c.elemtype, cas.elem, casePC(casi), chanrecvpc)
		}
		racenotify(c, c.recvx, nil)
	}
	if msanenabled && cas.elem != nil {
		msanwrite(cas.elem, c.elemtype.Size_)
	}
	if asanenabled && cas.elem != nil {
		asanwrite(cas.elem, c.elemtype.Size_)
	}
	recvOK = true
	qp = chanbuf(c, c.recvx)
	if cas.elem != nil {
		typedmemmove(c.elemtype, cas.elem, qp)
	}
	typedmemclr(c.elemtype, qp)
	c.recvx++
	if c.recvx == c.dataqsiz {
		c.recvx = 0
	}
	c.qcount--
	selunlock(scases, lockorder)
	goto retc

bufsend:
	// can send to buffer
	if raceenabled {
		racenotify(c, c.sendx, nil)
		raceReadObjectPC(c.elemtype, cas.elem, casePC(casi), chansendpc)
	}
	if msanenabled {
		msanread(cas.elem, c.elemtype.Size_)
	}
	if asanenabled {
		asanread(cas.elem, c.elemtype.Size_)
	}
	typedmemmove(c.elemtype, chanbuf(c, c.sendx), cas.elem)
	c.sendx++
	if c.sendx == c.dataqsiz {
		c.sendx = 0
	}
	c.qcount++
	selunlock(scases, lockorder)
	goto retc

recv:
	// can receive from sleeping sender (sg)
	recv(c, sg, cas.elem, func() { selunlock(scases, lockorder) }, 2)
	if debugSelect {
		print("syncrecv: cas0=", cas0, " c=", c, "\n")
	}
	recvOK = true
	goto retc

rclose:
	// read at end of closed channel
	selunlock(scases, lockorder)
	recvOK = false
	if cas.elem != nil {
		typedmemclr(c.elemtype, cas.elem)
	}
	if raceenabled {
		raceacquire(c.raceaddr())
	}
	goto retc

send:
	// can send to a sleeping receiver (sg)
	if raceenabled {
		raceReadObjectPC(c.elemtype, cas.elem, casePC(casi), chansendpc)
	}
	if msanenabled {
		msanread(cas.elem, c.elemtype.Size_)
	}
	if asanenabled {
		asanread(cas.elem, c.elemtype.Size_)
	}
	send(c, sg, cas.elem, func() { selunlock(scases, lockorder) }, 2)
	if debugSelect {
		print("syncsend: cas0=", cas0, " c=", c, "\n")
	}
	goto retc

retc:
	if caseReleaseTime > 0 {
		blockevent(caseReleaseTime-t0, 1)
	}
	return casi, recvOK

sclose:
	// send on closed channel
	selunlock(scases, lockorder)
	panic(plainError("send on closed channel"))
}

func (c *hchan) sortkey() uintptr {
	return uintptr(unsafe.Pointer(c))
}

// A runtimeSelect is a single case passed to rselect.
// This must match ../reflect/value.go:/runtimeSelect
type runtimeSelect struct {
	dir selectDir
	typ unsafe.Pointer // channel type (not used here)
	ch  *hchan         // channel
	val unsafe.Pointer // ptr to data (SendDir) or ptr to receive buffer (RecvDir)
}

// These values must match ../reflect/value.go:/SelectDir.
type selectDir int

const (
	_             selectDir = iota
	selectSend              // case Chan <- Send
	selectRecv              // case <-Chan:
	selectDefault           // default
)

//go:linkname reflect_rselect reflect.rselect
func reflect_rselect(cases []runtimeSelect) (int, bool) {
	if len(cases) == 0 {
		block()
	}
	sel := make([]scase, len(cases))
	orig := make([]int, len(cases))
	nsends, nrecvs := 0, 0
	dflt := -1
	for i, rc := range cases {
		var j int
		switch rc.dir {
		case selectDefault:
			dflt = i
			continue
		case selectSend:
			j = nsends
			nsends++
		case selectRecv:
			nrecvs++
			j = len(cases) - nrecvs
		}

		sel[j] = scase{c: rc.ch, elem: rc.val}
		orig[j] = i
	}

	// Only a default case.
	if nsends+nrecvs == 0 {
		return dflt, false
	}

	// Compact sel and orig if necessary.
	if nsends+nrecvs < len(cases) {
		copy(sel[nsends:], sel[len(cases)-nrecvs:])
		copy(orig[nsends:], orig[len(cases)-nrecvs:])
	}

	order := make([]uint16, 2*(nsends+nrecvs))
	var pc0 *uintptr
	if raceenabled {
		pcs := make([]uintptr, nsends+nrecvs)
		for i := range pcs {
			selectsetpc(&pcs[i])
		}
		pc0 = &pcs[0]
	}

	chosen, recvOK := selectgo(&sel[0], &order[0], pc0, nsends, nrecvs, dflt == -1)

	// Translate chosen back to caller's ordering.
	if chosen < 0 {
		chosen = dflt
	} else {
		chosen = orig[chosen]
	}
	return chosen, recvOK
}

func (q *waitq) dequeueSudoG(sgp *sudog) {
	x := sgp.prev
	y := sgp.next
	if x != nil {
		if y != nil {
			// middle of queue
			x.next = y
			y.prev = x
			sgp.next = nil
			sgp.prev = nil
			return
		}
		// end of queue
		x.next = nil
		q.last = x
		sgp.prev = nil
		return
	}
	if y != nil {
		// start of queue
		y.prev = nil
		q.first = y
		sgp.next = nil
		return
	}

	// x==y==nil. Either sgp is the only element in the queue,
	// or it has already been removed. Use q.first to disambiguate.
	if q.first == sgp {
		q.first = nil
		q.last = nil
	}
}
