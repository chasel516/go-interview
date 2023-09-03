package _chan

import (
	"sync/atomic"
	"unsafe"
)

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

// 等待队列是一个包含多个 sudog 结构体的链表，用于存储正在等待发送或接收数据的 goroutine。
// 当有数据可用时，等待队列中的 goroutine 会被唤醒并继续执行。
type waitq struct {
	first *sudog // 队列头部指针
	last  *sudog // 队列尾部指针
}

// sudog 结构体是一个用于等待队列中的 goroutine 的结构体，
// 它包含了等待的 goroutine 的信息，如等待的 channel、等待的元素值、
// 等待的方向（发送或接收）等。
type sudog struct {
	// 等待的 goroutine
	g *g

	// 指向下一个 sudog 结构体
	next *sudog

	// 指向上一个 sudog 结构体
	prev *sudog

	//等待队列的元素
	elem unsafe.Pointer

	// 获取锁的时间
	acquiretime int64

	// 释放锁的时间
	releasetime int64

	//用于实现自旋锁。当一个gorutine需要等待另一个gorutine操作完成，
	//而等待时间很短的情况下就会使用自旋锁。
	//它会先获取当前的ticket值，并将其加1。然后，它会不断地检查结构体中的ticket字段是否等于自己的ticket值，
	//如果相等就说明获取到了锁，否则就继续自旋等待。当锁被释放时，另一个goroutine会将ticket值加1，从而唤醒等待的goroutine。
	//需要注意的是，自旋锁适用于等待时间很短的场景，如果等待时间较长，就会造成CPU资源的浪费
	ticket uint32

	// 等待的 goroutine是否已经被唤醒
	isSelect bool

	//success 表示通道 c 上的通信是否成功。
	//如果 goroutine 是因为在通道 c 上接收到一个值而被唤醒，那么 success 为 true；
	//如果是因为通道 c 被关闭而被唤醒，那么 success 为 false。
	success bool

	//用于实现gorutine的堆栈转移
	//当一个 goroutine 调用另一个 goroutine 时，它会创建一个 sudog 结构体，并将自己的栈信息保存在 sudog 结构体的 parent 字段中。
	//然后，它会将 sudog 结构体加入到等待队列中，并等待被调用的 goroutine 执行完成。
	//当被调用的 goroutine 执行完成时，它会将 sudog 结构体从等待队列中移除，并将 parent 字段中保存的栈信息恢复到调用者的栈空间中。
	//这样，调用者就可以继续执行自己的任务了。
	//需要注意的是，sudog 结构体中的 parent 字段只在 goroutine 调用其他 goroutine 的时候才会被使用，
	//因此在普通的 goroutine 执行过程中，它是没有被使用的。
	parent *sudog // semaRoot binary tree

	//用于连接下一个等待的 sudog 结构体
	//等待队列是一个链表结构，每个 sudog 结构体都有一个 waitlink 字段，用于连接下一个等待的 sudog 结构体。
	//当被等待的 goroutine 执行完成时，它会从等待队列中移除对应的 sudog 结构体，
	//并将 sudog 结构体中的 waitlink 字段设置为 nil，从而将其从等待队列中移除。
	//需要注意的是，waitlink 字段只有在 sudog 结构体被加入到等待队列中时才会被使用。
	//在普通的 goroutine 执行过程中，waitlink 字段是没有被使用的。
	waitlink *sudog // g.waiting list or semaRoot

	//等待队列的尾部指针,waittail 字段指向等待队列的尾部 sudog 结构体。
	//当被等待的 goroutine 执行完成时，它会从等待队列中移除对应的 sudog 结构体，并将 sudog 结构体中的 waitlink 字段设置为 nil，
	//从而将其从等待队列中移除。同时，waittail 字段也会被更新为等待队列的新尾部。
	//需要注意的是，waittail 字段只有在 sudog 结构体被加入到等待队列中时才会被使用。
	//在普通的 goroutine 执行过程中，waittail 字段是没有被使用的。
	waittail *sudog // semaRoot

	//在golang中，goroutine是轻量级线程，其调度由golang运行时系统负责。当一个goroutine需要等待某些事件的发生时，
	//它可以通过阻塞等待的方式让出CPU资源，等待事件发生后再被唤醒继续执行。这种阻塞等待的机制是通过wait channel实现的。
	//在sudog结构体中，c字段指向的wait channel是一个用于等待某些事件发生的channel。
	//当一个goroutine需要等待某些事件时，它会创建一个sudog结构体，并将该结构体中的c字段指向wait channel。
	//然后，它会将该sudog结构体加入到wait channel的等待队列中，等待事件发生后再被唤醒继续执行。
	//当一个goroutine需要等待某些事件时，它会将自己加入到wait channel的等待队列中，并阻塞等待事件发生。
	//当事件发生后，wait channel会将等待队列中的goroutine全部唤醒，让它们继续执行。
	//这种机制可以有效地避免busy waiting，提高CPU利用率。
	c *hchan // channel
}

type lockRankStruct struct {
}

// Mutual exclusion locks.  In the uncontended case,
// as fast as spin locks (just a few user-level instructions),
// but on the contention path they sleep in the kernel.
// A zeroed Mutex is unlocked (no need to initialize each lock).
// Initialization is helpful for static lock ranking, but not required.
type mutex struct {
	// Empty struct-demo if lock ranking is disabled, otherwise includes the lock rank
	lockRankStruct
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	key uintptr
}

// Needs to be in sync with ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/reflectdata/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
// ../internal/reflectlite/type.go:/^type.rtype.
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}

type g struct {

	// goroutine 的栈信息
	stack       stack
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic *_panic // innermost panic - offset known to liblink
	_defer *_defer // innermost defer
	m      *m      // current m; offset known to arm liblink

	// goroutine 调度器上下文信息
	sched gobuf

	syscallsp uintptr // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc uintptr // if status==Gsyscall, syscallpc = sched.pc to use during gc
	stktopsp  uintptr // expected sp at top of stack, to check in traceback
	// param is a generic pointer parameter field used to pass
	// values in particular contexts where other storage for the
	// parameter would be difficult to find. It is currently used
	// in three ways:
	// 1. When a channel operation wakes up a blocked goroutine, it sets param to
	//    point to the sudog of the completed blocking operation.
	// 2. By gcAssistAlloc1 to signal back to its caller that the goroutine completed
	//    the GC cycle. It is unsafe to do so in any other way, because the goroutine's
	//    stack may have moved in the meantime.
	// 3. By debugCallWrap to pass parameters to a new goroutine because allocating a
	//    closure in the runtime is forbidden.
	param unsafe.Pointer

	// 原子级别的 goroutine 运行状态
	atomicstatus atomic.Uint32

	stackLock uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid      uint64
	schedlink guintptr

	waitsince int64 // approx time when the g become blocked

	// goroutine 等待的原因
	waitreason waitReason // if status==Gwaiting

	preempt       bool // preemption signal, duplicates stackguard0 = stackpreempt
	preemptStop   bool // transition to _Gpreempted on preemption; otherwise, just deschedule
	preemptShrink bool // shrink stack at synchronous safe point

	// asyncSafePoint is set if g is stopped at an asynchronous
	// safe point. This means there are frames on the stack
	// without precise pointer information.
	asyncSafePoint bool

	paniconfault bool // panic (instead of crash) on unexpected fault address
	gcscandone   bool // g has scanned stack; protected by _Gscan bit in status
	throwsplit   bool // must not split stack
	// activeStackChans indicates that there are unlocked channels
	// pointing into this goroutine's stack. If true, stack
	// copying needs to acquire channel locks to protect these
	// areas of the stack.
	activeStackChans bool
	// parkingOnChan indicates that the goroutine is about to
	// park on a chansend or chanrecv. Used to signal an unsafe point
	// for stack shrinking.
	parkingOnChan atomic.Bool

	raceignore     int8     // ignore race detection events
	sysblocktraced bool     // StartTrace has emitted EvGoInSyscall about this goroutine
	tracking       bool     // whether we're tracking this G for sched latency statistics
	trackingSeq    uint8    // used to decide whether to track this G
	trackingStamp  int64    // timestamp of when the G last started being tracked
	runnableTime   int64    // the amount of time spent runnable, cleared when running, only used when tracking
	sysexitticks   int64    // cputicks when syscall has returned (for tracing)
	traceseq       uint64   // trace event sequencer
	tracelastp     puintptr // last P emitted an event for this goroutine
	lockedm        muintptr

	// goroutine 等待的信号
	sig uint32

	writebuf   []byte
	sigcode0   uintptr
	sigcode1   uintptr
	sigpc      uintptr
	gopc       uintptr         // pc of go statement that created this goroutine
	ancestors  *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)
	startpc    uintptr         // pc of goroutine function
	racectx    uintptr
	waiting    *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt    []uintptr      // cgo traceback context
	labels     unsafe.Pointer // profiler labels
	timer      *timer         // cached timer for time.Sleep
	selectDone atomic.Uint32  // are we participating in a select and did someone win the race?

	// goroutineProfiled indicates the status of this goroutine's stack for the
	// current in-progress goroutine profile
	goroutineProfiled goroutineProfileStateHolder

	// Per-G GC state

	// gcAssistBytes is this G's GC assist credit in terms of
	// bytes allocated. If this is positive, then the G has credit
	// to allocate gcAssistBytes bytes without assisting. If this
	// is negative, then the G must correct this by performing
	// scan work. We track this in bytes to make it fast to update
	// and check for debt in the malloc hot path. The assist ratio
	// determines how this corresponds to scan work debt.
	gcAssistBytes int64
}
