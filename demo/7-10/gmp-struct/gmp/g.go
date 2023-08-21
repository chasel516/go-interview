package gmp

import (
	"sync/atomic"
	"unsafe"
)

// 取 goroutine 的首字母，主要保存 goroutine 的一些状态信息以及 CPU 的一些寄存器的值，
// 例如 IP 寄存器，以便在轮到本 goroutine 执行时，CPU 知道要从哪一条指令处开始执行。
// 当 goroutine 被调离 CPU 时，调度器负责把 CPU 寄存器的值保存在 g 对象的成员变量之中。
// 当 goroutine 被调度起来运行时，调度器又负责把 g 对象的成员变量所保存的寄存器值恢复到 CPU 的寄存器。
type g struct {
	// 堆栈参数,描述实际的堆栈内存,包括栈底和栈顶指针
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

	//syscallsp用于在 Goroutine 的状态为 Gsyscall（系统调用）时，指定在垃圾回收（GC）期间使用的 syscallsp 值。
	//在 Golang 的运行时系统中，当 Goroutine 进入系统调用时，它的状态会被设置为 Gsyscall。
	//系统调用是指 Goroutine 请求底层操作系统提供的服务，例如文件 I/O、网络 I/O等。
	//在进行垃圾回收期间，运行时系统需要执行一些特殊的操作，以确保垃圾回收器正常工作，并且不会意外地影响到系统调用过程。
	//为了实现这一点，运行时系统在 g 结构体中维护了 syscallsp 字段。
	//syscallsp 字段存储了一个地址值，表示在 GC 期间应使用的 syscallsp 值。
	//该值通常是指向系统调用的堆栈帧的栈指针（stack pointer）。
	//通过保存 syscallsp 值，垃圾回收器可以在 GC 过程中正确地恢复系统调用的状态。
	syscallsp uintptr // if status==Gsyscall, syscallsp = sched.sp to use during gc

	//用于在 Goroutine 的状态为 Gsyscall（系统调用）时，在垃圾回收（GC）期间使用的 syscallpc 值。
	//当 Goroutine 进入系统调用时，它的状态会被设置为 Gsyscall，表示该 Goroutine 正在执行一个系统调用操作。
	//跟上面的syscallsp字段类似，也是为了确保在进行垃圾回收期间，垃圾回收器的正常工作。
	//syscallpc的值通常是指向调用系统调用的指令（instruction）的程序计数器（program counter）。
	//也是为了垃圾回收器可以在 GC 过程中正确地恢复系统调用的执行状态。
	syscallpc uintptr // if status==Gsyscall, syscallpc = sched.pc to use during gc

	//用于在回溯（traceback）期间检查堆栈顶部的预期栈指针（stack pointer）。
	//回溯是一种在程序出现错误或异常情况时，追踪并打印调用堆栈信息的过程。这有助于开发人员定位问题的来源和上下文。
	//stktopsp 字段存储了一个地址值，表示预期的堆栈顶部的栈指针（stack pointer）。
	//在进行回溯时，运行时系统会检查堆栈的顶部，以确保其栈指针与 stktopsp 字段中的预期值匹配。
	//通过使用 stktopsp 字段，运行时系统可以在回溯期间验证堆栈的一致性，以确保栈指针的正确性。
	//这有助于检测堆栈溢出、错误的调用帧和其他潜在的问题。
	stktopsp uintptr // expected sp at top of stack, to check in traceback

	// param 是一个通用指针参数字段，用于传递 参数的其他存储空间将很难找到的情况下传递值。
	//目前有三种使用方式
	// 参数字段：
	// 1. 当一个通道操作唤醒一个被阻塞的 goroutine 时，它会将 param 设置为指向已完成阻塞操作的 sudog。
	//2. 通过 gcAssistAlloc1 向其调用者发出信号，表明该 goroutine 已完成GC 循环。以任何其他方式都是不安全的，因为 goroutine 的堆栈可能会在GC期间发生了移动；
	//3. 通过 debugCallWrap 向新的 goroutine 传递参数，因为在运行时分配闭包是被禁止的。
	param unsafe.Pointer

	//Goroutine的原子状态，用于表示Goroutine的状态，比如运行状态、等待状态、休眠状态等。
	atomicstatus atomic.Uint32

	//栈锁标志，用于在并发情况下对栈进行加锁
	stackLock uint32 // sigprof/scang lock; TODO: fold in to atomicstatus

	// Goroutine的唯一标识符，用于唯一标识每个Goroutine
	goid uint64

	//与调度器（scheduler）相关联的链表指针
	//调度器可以通过遍历 G 结构体的 schedlink 链表来选择下一个要执行的 Goroutine，从而实现调度和切换的功能。
	//通过使用 schedlink 字段，Golang 的运行时系统可以有效地管理和调度 Goroutine，实现并发执行和调度的功能。
	schedlink guintptr

	//表示 Goroutine（g 结构体）被阻塞的近似时间。
	//当 Goroutine 被阻塞时，runtime 系统会记录当前的时间，并将其存储在 waitsince 字段中。
	//通过使用 waitsince 字段，可以在调度器中追踪 Goroutine 的阻塞时间，用于调度和统计等用途。
	waitsince int64 // approx time when the g become blocked

	//表示 Goroutine（g 结构体）等待的原因。
	//waitreason 字段仅在 Goroutine 的状态（status 字段）为 Gwaiting（等待）时才会被使用。
	//它存储了 Goroutine 等待的具体原因，以便在调度器中进行跟踪和诊断。
	//waitreason 是一个枚举类型（enum），定义了多种可能的等待原因,比如：
	//waitReasonZero: 零值，表示无具体等待原因。
	//waitReasonIOWait: Goroutine 正在等待 I/O 操作完成。
	//waitReasonChannelSend: Goroutine 正在等待向通道发送数据的操作完成。
	//waitReasonChannelReceive: Goroutine 正在等待从通道接收数据的操作完成。
	//waitReasonSleep: Goroutine 正在等待休眠结束。
	waitreason waitReason // if status==Gwaiting

	//表示 Goroutine（g 结构体）是否收到了抢占信号。
	//当 preempt 字段为 true 时，表示该 Goroutine 已经收到了抢占信号。
	//抢占信号可以来自于其他更高优先级的 Goroutine 或运行时系统的调度器。
	//抢占信号的目的是中断当前 Goroutine 的执行，以便调度器可以切换到其他 Goroutine。
	//通过使用 preempt 字段，可以在合适的时机中断当前 Goroutine 的执行，以实现公平的 Goroutine 调度。
	//这对于防止某个 Goroutine 长时间占用 CPU 资源，从而导致其他 Goroutine 饥饿的情况是非常重要的。
	preempt bool // preemption signal, duplicates stackguard0 = stackpreempt

	//用于表示 Goroutine（g 结构体）在被抢占时的行为。当 preemptStop 字段为 true 时，表示该 Goroutine 在被抢占时会转变为 _Gpreempted 状态。
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

	// asyncSafePoint表示 Goroutine（g 结构体）是否在异步安全点（asynchronous safe point）被停止.
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

	//表示 Goroutine（g结构体）的堆栈是否已经被垃圾回收器扫描完成.
	//在进行垃圾回收时，垃圾回收器需要遍历并扫描 Goroutine 的堆栈，以确定堆栈上的对象是否可以被回收。
	//当 gcscandone 字段为 true 时，表示该 Goroutine 的堆栈已经被垃圾回收器完全扫描过了。
	//这意味着垃圾回收器已经检查了堆栈上的所有对象，并进行了必要的回收操作。
	//gcscandone 字段的设置受到 Goroutine 状态中的 _Gscan 位的保护。
	//_Gscan 位是 Goroutine 状态的一部分，用于标记 Goroutine 是否正在被垃圾回收器扫描。
	//gcscandone 字段的设置是受到 _Gscan 状态的保护的
	//这种保护机制确保了在垃圾回收过程中，只有在进行堆栈扫描的 Goroutine 才能更新 gcscandone 字段。
	//这样可以避免其他并发的 Goroutine 在未完成扫描的情况下访问或修改 gcscandone 字段，确保了数据的一致性
	gcscandone bool // g has scanned stack; protected by _Gscan bit in status

	//用于表示 Goroutine 是否禁止分割堆栈的标识.
	//当 throwsplit 字段为 true 时，表示该 Goroutine 被标记为禁止堆栈分割。堆栈分割是指在 Goroutine 运行时自动增长和收缩堆栈的过程。
	//Golang 的运行时系统为每个 Goroutine 分配一块固定大小的堆栈空间。
	//当 Goroutine 的堆栈空间不足以容纳当前的执行需求时，堆栈会自动增长，以满足更大的堆栈需求。
	//类似地，当 Goroutine 的堆栈空间超过一定的阈值时，堆栈会自动收缩，以释放不再需要的内存。
	//然而，在某些情况下，我们希望禁止 Goroutine 的堆栈分割,要求 Goroutine 在整个生命周期中保持固定大小的堆栈.
	//通过将 throwsplit 字段设置为 true，我们可以指示运行时系统不对该 Goroutine 进行堆栈的自动增长和收缩操作。
	//这可以确保 Goroutine 的堆栈始终保持固定的大小。
	throwsplit bool // must not split stack

	//表示是否存在指向该 Goroutine 堆栈的未锁定通道(channels)
	//当 activeStackChans 字段为 true 时，表示存在指向该 Goroutine 堆栈的未锁定通道。
	//换句话说，有其他 Goroutine 可能通过通道访问当前 Goroutine 的堆栈。
	//在 Golang 中，通道是一种用于 Goroutine 之间通信和同步的重要机制。
	//当一个 Goroutine 将值发送到通道或从通道接收值时，它会与其他 Goroutine 进行通信。这可能涉及对通道的锁定和解锁操作。
	//在某些情况下，通道可以指向其他 Goroutine 的堆栈，这意味着其他 Goroutine 可以通过通道访问当前 Goroutine 的堆栈数据。
	//当 activeStackChans 字段为 true 时，表示当前 Goroutine 存在这样的未锁定通道。
	//为了保护堆栈数据的完整性，当需要对当前 Goroutine 的堆栈进行复制（例如，在进行垃圾回收时），需要先获取通道的锁，
	//以防止其他 Goroutine 访问该堆栈。通过锁定通道，可以确保堆栈复制过程中的数据一致性。
	activeStackChans bool

	// 用于表示 Goroutine 即将在 chansend（通道发送）或 chanrecv（通道接收）操作上进行休眠。
	//它用于标记一个不安全点，以便进行堆栈缩减。
	//我们知道，通道是 Golang 中用于 Goroutine 之间进行通信的机制。
	//在某些情况下，当 Goroutine 需要在 chansend 或 chanrecv 操作上等待时，它会进入休眠状态，暂时停止执行，直到满足相应的条件。
	//当 Goroutine 即将在 chansend 或 chanrecv 操作上进行休眠时，parkingOnChan 字段被设置为 true。
	//这表示当前 Goroutine 正在等待通道操作完成，并准备进入休眠状态。
	//为了进行堆栈的收缩（stack shrinking），即释放不再需要的堆栈空间，需要在安全点（safe point）执行堆栈收缩操作。
	//但是，在 Goroutine 进入休眠状态之前，堆栈收缩是不安全的，因为收缩操作可能会干扰到正在进行的通道操作。
	//通过标记 parkingOnChan 字段为 true，可以将 Goroutine 的休眠操作作为不安全点进行标记。
	//在运行时系统的堆栈收缩机制中，会遇到这个标记，并在安全点之前等待 Goroutine 的休眠操作完成后再进行堆栈的收缩。
	parkingOnChan atomic.Bool

	//用于忽略竞争检测（race detection）事件。
	//竞争检测是Golang的一项重要功能，用于检测并发程序中的数据竞争情况，即多个goroutine同时访问和修改共享数据的情况。
	//竞争检测的目的是帮助开发者发现并修复潜在的并发问题。
	//我们可以使用go build/run时指定-race参数进行静态检测的调试，在go文件首行增加注释行 `// +build !race`可以告诉编译器不需要开启静态检测
	raceignore int8 // ignore race detection events

	//表示在当前 Goroutine（g 结构体）上是否已经发出了与系统调用（syscall）相关的跟踪事件.
	//标记当前 Goroutine 是否已经被追踪，并且与系统调用相关的事件已经被发出。
	//当 sysblocktraced 字段为 true 时，表示当前 Goroutine 已经被追踪，并且与系统调用相关的跟踪事件（EvGoInSyscall）已经被发出。
	//系统调用是指程序在执行期间通过操作系统提供的接口进行的一种操作，例如文件读写、网络通信等。
	//在跟踪和调试的过程中，记录 Goroutine 进入和离开系统调用的事件可以提供有关程序行为的详细信息。
	sysblocktraced bool // StartTrace has emitted EvGoInSyscall about this goroutine

	//用于指示是否正在跟踪该 Goroutine 的调度延迟统计信息。
	//调度延迟是指在多个 Goroutine 并发执行时，从一个 Goroutine 切换到另一个 Goroutine 的时间延迟。
	//Golang 的运行时系统具有用于跟踪和统计调度延迟的功能，以提供有关 Goroutine 的调度性能和行为的信息。
	//当 tracking 字段为 true 时，表示正在跟踪该 Goroutine 的调度延迟统计信息。
	//这意味着运行时系统会记录和统计与该 Goroutine 相关的调度延迟数据。
	//跟踪调度延迟可以帮助开发人员了解 Goroutine 的调度行为，识别潜在的性能问题和瓶颈，并进行性能调优。
	//通过收集和分析调度延迟统计信息，可以优化 Goroutine 的调度策略和资源利用，提高程序的并发性能。
	tracking bool // whether we're tracking this G for sched latency statistics

	//用于决定是否跟踪该 Goroutine的一个序列号.
	//在 Golang 的运行时系统中，跟踪（tracking）是一项资源密集型的操作，会占用额外的内存和计算资源。
	//为了优化性能，运行时系统不会对所有的 Goroutine 进行跟踪。相反，它使用 trackingSeq 字段来决定是否对该 Goroutine 进行跟踪。
	//trackingSeq 字段用于存储一个序列号（sequence number），该序列号用于标识该 Goroutine。
	//运行时系统会根据一些策略和算法，根据 trackingSeq 字段的值来决定是否跟踪该 Goroutine。
	//通过使用 trackingSeq 字段，运行时系统可以有选择地对一部分 Goroutine 进行跟踪，以平衡性能和资源的消耗。
	//对于没有跟踪的 Goroutine，它们的 trackingSeq 字段可能会被设置为一个特定的值（例如 0），表示不进行跟踪。
	trackingSeq uint8 // used to decide whether to track this G

	//用于记录最后一次开始跟踪该 Goroutine 的时间戳
	trackingStamp int64 // timestamp of when the G last started being tracked

	//用于记录 Goroutine 处于可运行状态的时间。该字段只在启用了 Goroutine 跟踪功能时使用.
	//Goroutine 的可运行状态是指它已经准备好并可以被调度执行的状态。
	//当 Goroutine 处于可运行状态时，它可能正在等待调度器选择执行它。
	//runnableTime 字段记录了 Goroutine 处于可运行状态的累计时间。
	//当 Goroutine 开始运行时，即被调度器选择执行时，runnableTime 字段会被清零。
	//然后，在 Goroutine 再次变为可运行状态之前，runnableTime 字段会记录 Goroutine 处于可运行状态的时间。
	//该字段的目的是帮助收集有关 Goroutine 的调度统计信息，例如计算 Goroutine 的可运行时间占据总运行时间的比例。
	//通过跟踪 Goroutine 的 runnableTime，可以了解 Goroutine 在调度器中的活跃程度和执行时间。
	runnableTime int64 // the amount of time spent runnable, cleared when running, only used when tracking

	//用于在跟踪中记录系统调用返回时的 CPU 时间戳。
	//当 Goroutine 执行系统调用并返回时，runtime 系统会记录 sysexitticks 字段的值，以标记系统调用返回时的 CPU 时间戳。
	//这个时间戳可以用于跟踪工具来分析系统调用的性能和行为。
	//通过记录系统调用返回的时间戳，可以分析系统调用的耗时以及与其他事件之间的时间关系。这对于性能调优和故障排查非常有用。
	sysexitticks int64 // cputicks when syscall has returned (for tracing)

	//用于跟踪事件的顺序号生成器。在 Golang 中，可以使用 trace 包来进行性能分析和跟踪应用程序的执行过程。
	//跟踪事件是指在应用程序的执行过程中发生的一系列重要事件，例如函数调用、系统调用、内存分配等。
	//这些事件可以帮助开发人员了解应用程序的执行流程、资源使用情况和性能瓶颈。
	//traceseq 字段用于生成跟踪事件的顺序号。
	//每当需要生成一个新的跟踪事件时，运行时系统会将 traceseq 字段的值作为该事件的顺序号，并将 traceseq 字段递增。
	//这样可以确保跟踪事件具有唯一的顺序号，以便在跟踪分析工具中进行准确的时间线重建和分析。
	traceseq uint64 // trace event sequencer

	//用于记录最后一个发出事件的 P（处理器）与该 Goroutine 之间的关联。
	//在 Golang 的运行时系统中，P 是一个处理器，用于执行 Goroutine。
	//当 Goroutine 需要执行时，调度器会将其分配给一个 P，以便在该处理器上执行。
	//tracelastp 字段用于跟踪与该 Goroutine 相关联的最后一个发出事件的 P。
	//这个字段的值是一个指向 P 的指针，指示最后一个与该 Goroutine 相关的事件发生的处理器。
	//跟踪与 Goroutine 相关的最后一个事件的 P 可以在调试和跟踪工具中使用。
	//通过记录最后一个事件的 P，可以了解 Goroutine 最后一次在哪个处理器上执行，有助于分析调度行为和性能特征。
	tracelastp puintptr // last P emitted an event for this goroutine

	//用于存储与当前 Goroutine 相关联的锁定的 m（调度器线程）。
	//当 Goroutine 成功获得一个锁时，lockedm 字段会被设置为持有该锁的 m 的指针。
	//这表示当前 Goroutine 正在持有特定的锁，并且与该锁关联的 m 负责执行该 Goroutine。
	//通过跟踪 lockedm 字段，运行时系统可以了解当前 Goroutine 所持有的锁和关联的调度器线程。
	//这对于锁的调度和资源的管理是至关重要的，确保并发访问的正确性和一致性。
	lockedm muintptr

	//用于存储与 Goroutine 相关的信号信息。
	//信号是在操作系统级别产生的事件或通知，可能包含中断、错误、终止等。
	//这些信号可以由操作系统、其他 Goroutine 或用户代码发送。
	//当一个 Goroutine 接收到信号时，runtime 系统会将相应的信号信息存储在 sig 字段中。
	//这样，Goroutine 可以在适当的时候对接收到的信号进行处理。
	sig uint32

	//用于存储用于写入数据的缓冲区。
	//writebuf 字段通常用于在 Goroutine 执行网络通信或文件操作时，临时存储要发送或写入的数据。
	//它提供了一个缓冲区，用于暂时保存待写入的字节数据，然后批量地发送或写入到相应的目标。
	//由于网络通信和文件操作涉及到读写效率和数据传输的优化，使用缓冲区可以减少系统调用的频率，提高数据传输的效率。
	//通过将数据先写入 writebuf 字段中的缓冲区，可以在一定程度上减少对底层资源的频繁访问和操作
	writebuf []byte

	//用于存储与 Goroutine 相关的信号代码（Signal Code）信息。
	//sigcode0 字段用于表示 Goroutine 接收到的信号的代码。信号代码是与特定信号相关联的标识符，用于描述信号的类型和原因。
	//当一个 Goroutine 接收到信号时，runtime 系统会将相应的信号代码存储在 sigcode0 字段中。
	//这样，Goroutine 可以根据信号代码来确定如何处理接收到的信号。
	sigcode0 uintptr

	//也是用于存储与 Goroutine 相关的信号代码（Signal Code）信息。
	//sigcode0 用于存储接收到的信号的主要代码，而 sigcode1 则用于存储与信号相关的附加信息或辅助代码.
	//sigcode0 字段通常用于表示主要的信号代码，用于标识接收到的信号的类型或操作。
	//例如，它可能用于表示 SIGSEGV（段错误）信号、SIGILL（非法指令）信号等。
	//而 sigcode1 字段则通常用于存储与信号相关的附加信息或辅助代码，用于提供更多关于信号的上下文或详细信息。
	//例如，它可能用于存储导致段错误的内存地址、非法指令的指令地址等。
	sigcode1 uintptr

	//用于存储与信号相关的程序计数器（Program Counter）.
	//程序计数器是一种特殊的寄存器，用于存储当前正在执行的指令的地址。它指示了正在执行的代码的位置，是处理器执行下一条指令的地址。
	//在信号处理的上下文中，当发生信号时，操作系统会中断正在执行的程序，并将控制权转移到信号处理程序中。
	//sigpc 字段用于存储与信号相关的程序计数器的值，即表示信号发生时的指令地址。
	//sigpc 字段的值可以用于追踪和记录信号发生的位置，帮助调试和诊断信号处理相关的问题。
	//通过记录程序计数器的值，可以知道在发生信号的时候程序执行到了哪个位置，从而帮助定位可能出现的错误或异常情况。
	sigpc uintptr

	//gopc用于存储该goroutine的go语句的程序计数器（program counter）值。
	//程序计数器是一个特殊的寄存器，用于存储当前正在执行的指令的地址。在Golang中，每个goroutine都有一个与之关联的程序计数器。
	//当一个新的goroutine被创建时，runtime系统会记录并存储该goroutine的创建点信息，即创建该goroutine的go语句的程序计数器值。
	//这个信息可以通过g结构体中的gopc字段来访问。
	//通过pc字段，我们可以追踪和识别每个goroutine的创建点，即调用go语句的位置。这对于调试和分析goroutine的创建和执行过程非常有用。
	//它可以帮助我们了解goroutine的起源和上下文，从而更好地理解程序的并发行为。
	//需要注意的是，gopcpc字段的值是一个指向具体指令的地址。
	//因为在Golang中的goroutine创建语句通常是go函数调用，所以pc字段的值可以理解为创建goroutine的go语句所在函数的程序计数器值。
	//通过使用pc字段，我们可以了解和跟踪goroutine的创建点，从而更好地理解程序的执行和并发模式。
	//这对于调试、性能优化和并发问题排查都非常有帮助。
	gopc uintptr // pc of go statement that created this goroutine

	//用于存储创建当前 Goroutine 的祖先 Goroutine 的信息.
	//例如祖先 Goroutine 的 ID、栈跟踪信息等。
	//ancestors 字段在以下情况下被使用：
	//1.只有在启用了 debug.tracebackancestors 标志时才会被使用。
	//2. 当需要追踪和记录创建当前 Goroutine 的祖先 Goroutine 信息时，运行时系统会将祖先 Goroutine 的相关信息存储在 ancestors 字段指向的切片中。
	//debug.tracebackancestors 是一个调试标志，用于启用在 Goroutine 栈跟踪信息中记录祖先 Goroutine 信息的功能。
	//当该标志启用时，运行时系统会在 Goroutine 的栈跟踪信息中包含创建该 Goroutine 的祖先 Goroutine 的信息。
	//通过使用 ancestors 字段和 debug.tracebackancestors 标志，Golang 的运行时系统可以提供更丰富的 Goroutine 栈跟踪信息，
	//包括祖先 Goroutine 的相关信息。这有助于调试和定位 Goroutine 的创建和调用关系，特别是在复杂的并发程序中。
	ancestors *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)

	//用于存储Goroutine 函数的起始指令的地址.在 Golang 中，每个 Goroutine 都与一个特定的函数关联，该函数定义了 Goroutine 的入口点和要执行的代码。
	//startpc 字段记录了与该 Goroutine 关联的函数的起始指令地址。
	//通过 startpc 字段，可以得到 Goroutine 函数的起始位置，也就是 Goroutine 执行的起始点。
	//这对于追踪和调试 Goroutine 的执行流程和代码路径非常有用。
	//可以根据 startpc 字段的值，定位到 Goroutine 执行的具体代码位置。
	startpc uintptr // pc of goroutine function

	//用于存储与竞态检测（race detection）相关的上下文信息。
	//竞态检测是一种用于检测并发程序中可能存在的数据竞态（data race）问题的技术。
	//数据竞态是指多个 Goroutine 同时访问共享数据，并且至少有一个访问是写操作，而且这些访问没有适当的同步操作。
	//竞态检测的目标是在程序运行时检测和报告这些潜在的并发错误。
	//racectx 字段用于存储与竞态检测相关的上下文信息。
	//这些上下文信息可能包括跟踪竞态检测的状态、报告错误的位置、数据访问的历史等。
	//具体的上下文信息和使用方式取决于竞态检测工具和运行时配置。
	racectx uintptr

	// waiting 字段是一个指向 sudog 结构体的指针，用于存储当前 Goroutine 等待的 sudog 结构体。
	//sudog（也称为唤醒原语，synchronization descriptor）是一种在 Golang 的运行时系统中用于实现 Goroutine 的等待和唤醒机制的数据结构。
	//当一个 Goroutine 需要等待某个条件满足时，它会进入等待状态，并将相关的 sudog 结构体添加到 waiting 字段中。
	//waiting 字段存储了当前 Goroutine 正在等待的 sudog 结构体的指针。
	//这些 sudog 结构体通常与锁、通道或其他同步原语相关联，并用于在特定条件满足时唤醒 Goroutine。
	//waiting 字段中的 sudog 结构体按照特定的锁顺序进行排列，以确保等待和唤醒的正确顺序。
	//当条件满足时，其他 Goroutine 可能会通过唤醒等待队列中的 sudog 结构体来通知等待中的 Goroutine 继续执行。
	waiting *sudog // sudog structures this g is waiting on (that have a valid elem ptr); in lock order

	//用于存储与 cgo相关的回溯（traceback）上下文信息.
	//在 Golang 中，cgo 是一种机制，允许在 Go 代码中调用 C 语言代码或使用 C 语言库。
	//当使用 cgo 进行混合编程时，Golang 的运行时系统可以跟踪与 cgo 相关的调用链，以便在发生问题或错误时进行回溯和调试。
	//cgoCtxt 字段是一个 uintptr 类型的切片，其中的每个元素表示 cgo 调用链中的一个函数的程序计数器（Program Counter）值。
	//程序计数器表示当前正在执行的指令的地址，通过记录 cgo 调用链中每个函数的程序计数器，可以构建完整的回溯上下文，
	//帮助定位和追踪与 cgo 相关的问题。
	cgoCtxt []uintptr // cgo traceback context

	//用于存储与性能分析器（profiler）相关的标签信息.这些标签可以用于标识和组织性能数据
	labels unsafe.Pointer // profiler labels

	//用于缓存 time.Sleep 操作所使用的定时器。
	//time.Sleep 是 Golang 中用于暂停执行一段时间的函数。它常用于控制 Goroutine 的执行间隔或实现延迟操作。
	//为了避免每次调用 time.Sleep 都创建新的定时器，runtime 使用 timer 字段来缓存一个定时器，以便重复使用。
	//timer 结构体包含了与定时器相关的信息，如过期时间、触发时间等。
	//通过缓存一个定时器，可以避免在每次调用 time.Sleep 时都进行定时器的创建和销毁，从而提高性能和效率。
	timer *timer // cached timer for time.Sleep

	//用于表示当前 Goroutine是否参与了 select 操作，并且是否存在竞争的情况
	//在 Golang 中，select 语句用于在多个通道操作之间进行非阻塞选择。
	//当一个 Goroutine执行 select 语句时，它会尝试在多个通道上进行操作，并等待其中一个操作完成。
	//如果多个 Goroutine同时执行 select，并且某个操作同时满足多个 Goroutine 的条件，就会发生竞争。
	//selectDone 字段通过使用 atomic.Uint32 类型的原子操作，表示当前 Goroutine是否参与了 select 操作，并且是否存在竞争的情况。
	//如果 selectDone 字段的值为非零（1），表示当前 Goroutine已经参与了 select 操作，并且可能存在竞争的情况。
	//竞争条件可能导致多个 Goroutine同时选择同一个操作，进而引发不确定的行为。
	//通过使用 selectDone 字段，可以标记当前 Goroutine参与了 select 操作，并通知其他 Goroutine不再进行竞争，从而避免竞争问题。
	selectDone atomic.Uint32 // are we participating in a select and did someone win the race?

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

// 描述栈的数据结构，栈的范围：[lo, hi)
type stack struct {
	// 栈顶，低地址
	lo uintptr
	// 栈低，高地址
	hi uintptr
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
