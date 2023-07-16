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

	//Goroutine的运行状态和上下文信息。(下面有gobuf结构体各个字段的注释)
	sched gobuf

	//syscallsp，syscallpc用于系统调用的堆栈指针和程序计数器
	syscallsp uintptr // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc uintptr // if status==Gsyscall, syscallpc = sched.pc to use during gc

	//Goroutine的堆栈边界
	stktopsp uintptr // expected sp at top of stack, to check in traceback

	// param字段是一个unsafe.Pointer类型的指针，用于传递参数。它是一个通用的字段，可以由用户自定义使用。
	//param字段的目的是为了在Goroutine的生命周期中传递任意类型的参数。
	//在创建Goroutine时，可以将需要传递的参数存储到param字段中，然后在Goroutine内部通过类型转换来获取参数的具体值。
	//由于param字段的类型是unsafe.Pointer，它可以指向任意类型的指针，包括指针类型和非指针类型。
	//这使得它具有很高的灵活性，可以适应各种参数类型的传递需求。
	//具体用法见：https://gitee.com/phper95/go-interview/blob/master/demo/7-10/gmp-struct/main.go
	param unsafe.Pointer

	//Goroutine的原子状态，用于表示Goroutine的状态和属性
	atomicstatus atomic.Uint32

	//栈锁标志，用于在并发情况下对栈进行加锁
	stackLock uint32 // sigprof/scang lock; TODO: fold in to atomicstatus

	// Goroutine的唯一标识符，用于唯一标识每个Goroutine
	goid uint64

	//与调度器（scheduler）相关联的链表指针
	//调度器可以通过遍历 G 结构体的 schedlink 链表来选择下一个要执行的 Goroutine，从而实现调度和切换的功能。
	//需要注意的是，schedlink 字段是在运行时系统内部使用的，对于用户代码是不可见的。它是用于调度器和调度算法的内部数据结构。
	//通过使用 schedlink 字段，Golang 的运行时系统可以有效地管理和调度 Goroutine，实现并发执行和调度的功能。
	schedlink guintptr

	//表示 Goroutine（G 结构体）被阻塞的近似时间。
	//当 Goroutine 被阻塞时，runtime 系统会记录当前的时间，并将其存储在 waitsince 字段中。
	//通过使用 waitsince 字段，可以在调度器中追踪 Goroutine 的阻塞时间，用于调度和统计等用途。
	waitsince int64 // approx time when the g become blocked

	//表示 Goroutine（G 结构体）等待的原因。
	//waitreason 字段仅在 Goroutine 的状态（status 字段）为 Gwaiting（等待）时才会被使用。
	//它存储了 Goroutine 等待的具体原因，以便在调度器中进行跟踪和诊断。
	//waitreason 是一个枚举类型（enum），定义了多种可能的等待原因,比如：
	//waitReasonZero: 零值，表示无具体等待原因。
	//waitReasonIOWait: Goroutine 正在等待 I/O 操作完成。
	//waitReasonChannelSend: Goroutine 正在等待向通道发送数据的操作完成。
	//waitReasonChannelReceive: Goroutine 正在等待从通道接收数据的操作完成。
	//waitReasonSleep: Goroutine 正在等待休眠结束。
	waitreason waitReason // if status==Gwaiting

	//表示 Goroutine（G 结构体）是否收到了抢占信号。
	//当 preempt 字段为 true 时，表示该 Goroutine 已经收到了抢占信号。
	//抢占信号可以来自于其他更高优先级的 Goroutine 或运行时系统的调度器。
	//抢占信号的目的是中断当前 Goroutine 的执行，以便调度器可以切换到其他 Goroutine。
	//通过使用 preempt 字段，可以在合适的时机中断当前 Goroutine 的执行，以实现公平的 Goroutine 调度。
	//这对于防止某个 Goroutine 长时间占用 CPU 资源，从而导致其他 Goroutine 饥饿的情况是非常重要的。
	preempt bool // preemption signal, duplicates stackguard0 = stackpreempt

	//用于表示 Goroutine（G 结构体）在被抢占时的行为。当 preemptStop 字段为 true 时，表示该 Goroutine 在被抢占时会转变为 _Gpreempted 状态。
	//_Gpreempted 是一个特殊的 Goroutine 状态，表示该 Goroutine 已被抢占。
	//相反，当 preemptStop 字段为 false 时，表示该 Goroutine 在被抢占时只会被调度器暂时挂起（deschedule），而不会转变为 _Gpreempted 状态。
	//_Gpreempted 状态的 Goroutine 会在恢复执行时立即被中断，以便让调度器可以重新调度它。
	//这样可以确保在 Goroutine 被抢占后能够尽快恢复到调度器的控制下。
	//preemptStop 字段的设置在一些特定的调度和抢占策略中是有意义的。
	//例如，在 M:N 调度模型中，当 Goroutine 被抢占时，它可能会在另一个 M 上继续执行。
	//设置 preemptStop 字段为 true 可以确保被抢占的 Goroutine 被中断并转变为 _Gpreempted 状态，以便在其他 M 上重新调度。
	preemptStop bool // transition to _Gpreempted on preemption; otherwise, just deschedule

	//用于表示在同步安全点（synchronous safe point）时是否缩小（shrink） Goroutine 的堆栈.
	//同步安全点是指在 Goroutine 执行期间的某个特定位置，它是一个安全的点，可以进行堆栈操作而不会影响 Goroutine 的一致性。
	//在同步安全点，Goroutine 的栈可以进行调整和缩小。
	//当 preemptShrink 字段为 true 时，表示该 Goroutine 在同步安全点时可以进行堆栈的缩小操作。
	//堆栈缩小可以释放不再使用的堆栈空间，从而减少内存的消耗。
	//堆栈的缩小操作通常是由运行时系统自动触发的，而 preemptShrink 字段可以用于控制该操作的行为。
	//在同步安全点时，如果 preemptShrink 字段为 true，则运行时系统会考虑对 Goroutine 的堆栈进行缩小操作。
	preemptShrink bool // shrink stack at synchronous safe point

	// asyncSafePoint表示 Goroutine（G 结构体）是否在异步安全点（asynchronous safe point）被停止.
	//当 asyncSafePoint 字段为 true 时，表示该 Goroutine 在异步安全点被停止。
	//异步安全点是 Goroutine 停止的一种特殊状态，可能涉及到对帧的扫描和处理，例如垃圾回收器的工作。
	//在异步安全点停止的情况下，Goroutine 的堆栈上可能存在一些帧，这些帧上的指针信息可能不是精确的,可能会对某些运行时操作和调试工具产生影响。
	asyncSafePoint bool

	//用于表示在出现意外的故障地址时，是否触发 panic（而不是崩溃）
	//注意panic 和crash这两个概念在golang中是有区别的
	//panic是一种可被程序捕获和处理的运行时错误，它会导致程序从当前函数传播到上层函数，并执行延迟函数。
	//而 crash 是指程序遇到无法恢复的严重错误，导致程序立即终止，并触发操作系统级别的错误处理机制。

	//在正常情况下，当程序访问到一个无效或未映射的内存地址时，运行时系统会引发一个崩溃（crash），导致程序终止。
	//这种崩溃是为了保护程序免受无效内存访问的影响。
	//但是，当 paniconfault 字段设置为 true 时，表示当程序遇到意外的故障地址时，会触发一个 panic，而不是直接崩溃。
	//通常情况下，paniconfault 字段不应该被设置为 true，除非你需要在出现意外的故障地址时进行特殊处理，
	//例如在调试或诊断过程中,设置为 true 可以捕获并处理无效内存访问的情况，使程序能够继续执行。
	paniconfault bool // panic (instead of crash) on unexpected fault address

	gcscandone bool // g has scanned stack; protected by _Gscan bit in status
	throwsplit bool // must not split stack
	// activeStackChans indicates that there are unlocked channels
	// pointing into this goroutine's stack. If true, stack
	// copying needs to acquire channel locks to protect these
	// areas of the stack.
	activeStackChans bool
	// parkingOnChan indicates that the goroutine is about to
	// park on a chansend or chanrecv. Used to signal an unsafe point
	// for stack shrinking.
	parkingOnChan atomic.Bool

	//用于忽略竞争检测（race detection）事件。
	//竞争检测是Golang的一项重要功能，用于检测并发程序中的数据竞争情况，即多个goroutine同时访问和修改共享数据的情况。
	//竞争检测的目的是帮助开发者发现并修复潜在的并发问题。
	//我们可以使用go build/run时指定-race参数进行静态检测的调试，在go文件首行增加注释行 `// +build !race`可以告诉编译器不需要开启静态检测
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

	//该goroutine的go语句的程序计数器（program counter）值。
	//程序计数器是一个特殊的寄存器，用于存储当前正在执行的指令的地址。在Golang中，每个goroutine都有一个与之关联的程序计数器。
	//当一个新的goroutine被创建时，runtime系统会记录并存储该goroutine的创建点信息，即创建该goroutine的go语句的程序计数器值。
	//这个信息可以通过g结构体中的gopc字段来访问。
	//通过pc字段，我们可以追踪和识别每个goroutine的创建点，即调用go语句的位置。这对于调试和分析goroutine的创建和执行过程非常有用。
	//它可以帮助我们了解goroutine的起源和上下文，从而更好地理解程序的并发行为。
	//需要注意的是，gopcpc字段的值是一个指向具体指令的地址。
	//因为在Golang中的goroutine创建语句通常是go函数调用，所以pc字段的值可以理解为创建goroutine的go语句所在函数的程序计数器值。
	//通过使用pc字段，我们可以了解和跟踪goroutine的创建点，从而更好地理解程序的执行和并发模式。
	//这对于调试、性能优化和并发问题排查都非常有帮助。
	gopc       uintptr         // pc of go statement that created this goroutine
	ancestors  *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)
	startpc    uintptr         // pc of goroutine function
	racectx    uintptr
	waiting    *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt    []uintptr      // cgo traceback context
	labels     unsafe.Pointer // profiler labels
	timer      *timer         // cached timer for time.Sleep
	selectDone atomic.Uint32  // are we participating in a select and did someone win the race?

	//在Golang的runtime包中，g结构体中的goroutineProfiled字段用于指示当前正在进行的goroutine分析（goroutine profiling）的堆栈状态。
	//goroutine profiling是一种性能分析工具，用于收集和分析正在运行的goroutine的信息，以了解其调用关系、执行时间和内存消耗等方面的情况。
	//goroutineProfiled字段用于跟踪和记录这种分析的状态。
	//具体来说，goroutineProfiled字段的值表示该goroutine的堆栈状态是否正在进行分析。
	//如果goroutineProfiled为true，则表示该goroutine的堆栈正在进行分析，即正在进行goroutine profiling。
	//反之，如果goroutineProfiled为false，则表示该goroutine的堆栈不在当前进行分析。
	//通过将goroutineProfiled字段与其他分析工具相关的字段进行配合使用，可以在分析过程中选择性地收集和分析特定的goroutine。
	//这有助于精确控制分析的范围和粒度，以便更好地理解和优化程序的性能。
	//需要注意的是，goroutineProfiled字段是在运行时系统内部维护的，并且对用户代码是不可见的。它是用于内部性能分析和调试工具的一部分。
	//通过使用goroutineProfiled字段，Golang的运行时系统能够准确地追踪和记录正在进行goroutine profiling的goroutine堆栈的状态，以提供更有效的性能分析和优化手段。
	goroutineProfiled goroutineProfileStateHolder

	//gcAssistBytes字段的值为正数时，表示Goroutine具有一定的垃圾回收辅助信用，可以在不进行辅助工作的情况下继续分配内存。
	//当Goroutine分配内存时，gcAssistBytes字段会递减。
	//而当gcAssistBytes字段的值为负数时，表示Goroutine需要通过执行扫描工作来纠正负债。
	//这意味着Goroutine需要帮助进行垃圾回收，以还清负债。
	//gcAssistBytes字段的设计考虑了性能的因素。通过以字节数为单位进行计算和跟踪，可以更快速地更新和检查是否存在垃圾回收负债。
	//而gcAssistBytes字段与扫描工作负债之间的对应关系是由辅助比例（assist ratio）来确定的。
	//辅助比例决定了gcAssistBytes字段与扫描工作负债之间的转换关系。
	//需要注意的是，gcAssistBytes字段是在运行时系统内部维护的，并且对用户代码是不可见的。
	//它是用于内部垃圾回收调度和性能优化的一部分。
	//通过使用gcAssistBytes字段，Golang的运行时系统可以根据Goroutine的内存分配情况来自动调整垃圾回收辅助工作的执行。
	//这样可以有效地管理内存和垃圾回收的负载，以提高程序的执行效率。
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
