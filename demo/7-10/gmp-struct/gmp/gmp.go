package gmp

import (
	"sync/atomic"
	"unsafe"
)

type g struct {
	// 堆栈参数,描述实际的堆栈内存,包括栈底和栈顶指针等
	stack stack // offset known to runtime/cgo

	// 用于栈的扩张和收缩检查，抢占标志,用于检测栈溢出情况
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	// 指向当前Goroutine的panic信息，用于处理异常
	_panic *_panic // innermost panic - offset known to liblink

	//指向当前Goroutine的延迟调用（defer）信息，用于处理延迟函数的执行
	_defer *_defer // innermost defer

	//指向执行该Goroutine的工作线程（M）的指针，用于Goroutine的调度和执行
	m *m // current m; offset known to arm liblink

	//Goroutine的运行状态和上下文信息。
	sched gobuf

	//syscallsp，syscallpc用于系统调用的堆栈指针和程序计数器
	syscallsp uintptr // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc uintptr // if status==Gsyscall, syscallpc = sched.pc to use during gc

	//Goroutine的堆栈边界
	stktopsp uintptr // expected sp at top of stack, to check in traceback

	// param字段是一个unsafe.Pointer类型的指针，用于传递参数。它是一个通用的字段，可以由用户自定义使用。param字段的目的是为了在Goroutine的生命周期中传递任意类型的参数。在创建Goroutine时，可以将需要传递的参数存储到param字段中，然后在Goroutine内部通过类型转换来获取参数的具体值。由于param字段的类型是unsafe.Pointer，它可以指向任意类型的指针，包括指针类型和非指针类型。这使得它具有很高的灵活性，可以适应各种参数类型的传递需求。具体用法见：https://gitee.com/phper95/go-interview/blob/master/demo/7-10/gmp-struct/main.go
	param unsafe.Pointer

	//Goroutine的原子状态，用于表示Goroutine的状态和属性
	atomicstatus atomic.Uint32

	stackLock  uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid       uint64
	schedlink  guintptr
	waitsince  int64      // approx time when the g become blocked
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
	sig            uint32
	writebuf       []byte
	sigcode0       uintptr
	sigcode1       uintptr
	sigpc          uintptr
	gopc           uintptr         // pc of go statement that created this goroutine
	ancestors      *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)
	startpc        uintptr         // pc of goroutine function
	racectx        uintptr
	waiting        *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt        []uintptr      // cgo traceback context
	labels         unsafe.Pointer // profiler labels
	timer          *timer         // cached timer for time.Sleep
	selectDone     atomic.Uint32  // are we participating in a select and did someone win the race?

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

type gobuf struct {

	//堆栈指针（Stack Pointer），表示Goroutine当前的堆栈位置
	sp uintptr // 堆栈指针

	//程序计数器(Program Counter)，表示Goroutine当前执行的指令地址
	pc uintptr // 程序计数器

	//当前Goroutine的g结构体指针，用于快速访问Goroutine的其他信息
	g guintptr // 当前Goroutine的g结构体指针

	// 保存的上下文指针，用于切换Goroutine时保存和恢复Goroutine的上下文信息
	ctxt unsafe.Pointer

	// 返回值（Return Value），用于保存函数调用返回时的返回值
	ret sys.Uintreg

	// 返回地址（Link Register），表示函数调用返回时需要跳转的地址
	lr uintptr

	//基指针(Base Pointer)，用于指示当前函数的栈帧
	bp uintptr
}
