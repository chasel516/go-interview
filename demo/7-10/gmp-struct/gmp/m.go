package gmp

import (
	"sync/atomic"
	"unsafe"
)

// m 代表工作线程，保存了自身使用的栈信息
type m struct {
	//g0 字段存储了具有调度堆栈的 Goroutine 的信息。g0被称为工作线程（也叫内核线程）。
	// 执行用户 goroutine 代码时，使用用户 goroutine 自己的栈，因此调度时会发生栈的切换
	//它是一个指向 g 结构体的指针，该结构体包含了 Goroutine 的状态、堆栈指针、程序计数器等信息。
	//g0 是一个特殊的 Goroutine，它用于管理调度和协程的创建、销毁和切换。
	//通过使用 g0 字段，运行时系统可以管理和调度具有调度堆栈的 Goroutine。
	//它包含了关键的上下文信息，使得 Goroutine 能够在 M（线程）上正确地执行和切换
	g0 *g // goroutine with scheduling stack

	//morebuf用于存储在执行 "morestack" 操作时传递给更大栈的参数。
	//"morestack" 是一个内部函数，用于处理 Goroutine 的栈扩展操作。
	//当 Goroutine 的栈空间不足以容纳当前的执行需求时，会触发 "morestack" 操作来分配更大的栈空间。
	//morebuf 字段用于保存传递给 "morestack" 操作的参数 gobuf。
	//前面介绍过，gobuf 是一个结构体，用于描述 Goroutine 的栈和寄存器等信息。
	//通过将 gobuf 参数存储在 morebuf 字段中，可以在执行 "morestack" 操作时提供必要的参数。
	//这样，"morestack" 函数就能够了解当前 Goroutine 的栈和寄存器状态，并进行栈扩展操作。
	morebuf gobuf // gobuf arg to morestack

	//用于在 ARM 架构中进行除法和取模运算的分母.
	//在 ARM 架构中，整数除法和取模运算（即除法的余数计算）通常是使用特定的分母进行计算，以提高性能和效率。
	//divmod 字段用于存储在进行这些运算时使用的分母值。
	//这个值是在编译阶段由编译器和链接器（liblink）确定，并且对于特定的 ARM 架构是已知的。
	divmod uint32 // div/mod denominator for arm - known to liblink

	//用于对齐下一个字段的内存布局，使其在 8 字节边界上对齐
	//在某些架构中，对于性能和内存访问的优化，数据在内存中需要按照特定的对齐方式存储。
	//例如，有些处理器要求 8 字节对齐，即某个数据的地址必须是 8 的倍数。
	//为了确保下一个字段在 8 字节边界上对齐，_ 字段被插入到结构体中作为填充字段。
	//该字段本身不存储任何实际的数据，它只是占用一定的内存空间，以使下一个字段在对齐位置上对齐。
	_ uint32 // align next field to 8 bytes

	// 用于提供给调试器使用的进程标识符。
	//procid 字段主要是为了调试器能够识别和跟踪与特定 M（Machine）关联的进程。
	//调试器是开发和调试程序时常用的工具，它可以与运行中的程序交互并提供调试信息。
	//procid 字段的具体值可以用于唯一标识与 M 关联的进程，并提供给调试器使用。
	//通过 procid 字段，调试器可以根据这个标识符来定位和跟踪特定的进程，以便进行调试操作和收集调试信息。
	procid uint64 // for debuggers, but offset not hard-coded

	// gsignal 字段是一个指向 g 结构体的指针，用于表示处理信号的 Goroutine。
	//在操作系统中，信号是用于通知进程发生了某个事件或需要进行某个操作的一种异步通知机制。
	//Golang 运行时系统中的 gsignal 字段用于保存负责处理信号的 Goroutine 的信息。
	//gsignal 字段存储了一个特殊的 Goroutine，它负责处理与运行时系统相关的信号。
	//当发生信号时，运行时系统会将信号传递给 gsignal 所指定的 Goroutine，并由该 Goroutine 进行相应的信号处理操作。
	//通过使用 gsignal 字段，Golang 的运行时系统可以将信号处理的责任委托给特定的 Goroutine，以便进行信号处理和相关操作。
	//这有助于将信号处理与其他 Goroutine 的执行分离，提高信号处理的可靠性和响应性。
	gsignal *g // signal-handling g

	//用于表示为信号处理分配的栈空间.
	//在 Golang 的运行时系统中，处理信号的操作通常需要在一个独立的栈上执行，以保证信号处理的可靠性和安全性。
	//在Golang的runtime包中，gsignal stack（信号堆栈）是用来处理操作系统信号的栈空间。
	//当操作系统发送信号给Go程序时，信号处理函数会在gsignal stack上运行。
	//这个栈是独立于普通的goroutine栈的，用于专门处理信号的相关操作。
	//gsignal stack的主要作用是提供一个独立的执行环境，确保信号处理函数能够正常运行而不受其他goroutine的影响。
	//在处理信号期间，runtime会禁止抢占和栈扩展，以确保信号处理函数的运行不会被干扰。
	//由于信号处理函数需要尽可能地简洁和高效，gsignal stack的大小是固定的，并且相对较小。
	//这是因为在信号处理期间，只能执行少量的操作，例如发送或接收信号、终止程序等。过多的操作可能会带来不可预知的问题。
	//需要注意的是，gsignal stack不同于goroutine栈，它是专门用于处理信号的，而goroutine栈则用于正常的程序执行。
	//这样的设计可以有效地隔离信号处理函数和普通程序逻辑，提高信号处理的可靠性和安全性。
	goSigStack gsignalStack // Go-allocated signal handling stack

	//表示 M 的信号掩码，用于控制哪些信号可以中断 M 的执行
	sigmask sigset // storage for saved signal mask

	//用于存储线程本地存储（TLS）的数据。
	//TLS 是一种机制，用于为每个线程提供独立的存储空间，使得每个线程都可以在其中存储和访问自己的数据， 而不会与其他线程的数据发生冲突。
	//tls 字段用于存储 TLS 的数据，以便每个线程可以访问自己的 TLS 数据。
	//在 m 结构体中，tls 字段是一个固定大小的数组，其中的元素类型为 uintptr。
	//每个元素可以存储一个指针或整数值，用于表示线程特定的数据。
	//通过使用 tls 字段，Golang 的运行时系统为每个线程提供了一个独立的存储空间，可以在其中存储线程特定的数据。
	//这有助于实现线程间的数据隔离和线程安全性。
	tls [tlsSlots]uintptr // thread-local storage (for x86 extern register)

	//M 的启动函数，表示在新创建的线程上运行的起始函数。在线程创建时，将执行此函数。
	//mstartfn 字段指向一个 func() 类型的函数，该函数作为 M 启动时的入口点。
	//它定义了 M 在启动时要执行的操作，通常是执行一些初始化工作，然后开始调度和执行 Goroutine
	mstartfn func()

	//表示当前正在执行的 Goroutine。curg 是一个指向 g 结构体的指针。在 M 上运行的 Goroutine 会存储在 curg 中。
	curg *g // current running goroutine

	//用于表示在发生致命信号（fatal signal）时正在执行的 Goroutine。
	//ghtsig 字段的值是一个 guintptr 类型的整数，通常是一个指针值。
	//它指向正在执行的 Goroutine 。
	//通过使用 caughtsig 字段，Golang 的运行时系统可以在发生致命信号时记录正在执行的 Goroutine，以便进行相应的处理和调试。
	//这有助于在程序崩溃时确定造成崩溃的 Goroutine。
	caughtsig guintptr // goroutine running during fatal signal

	//用于表示与当前 M（Machine）相关联的 P（Processor）。
	//在 Golang 的并发模型中，P 是调度器（Scheduler）用于管理 Goroutine 的上下文的单位。
	//每个 P 都维护了一组 Goroutine 队列和相关的调度状态。p 字段用于存储与当前 M 相关联的 P 的地址。
	//puintptr 是一个 uintptr 类型的别名，它用于保存指向 P 的地址。
	//通过使用 p 字段，当前 M 可以知道它所关联的 P，并在需要时与该 P 进行交互。
	//如果 M 正在执行 Go 代码，则 p 字段将保存与之关联的 P 的地址。
	//如果 M 当前不在执行 Go 代码（例如在系统调用中），则 p 字段将为 nil。
	//通过使用 p 字段，Golang 的运行时系统可以将 M 与正确的 P 关联起来，
	//并确保 M 在执行 Go 代码时可以正确地与所属的 P 进行通信和调度。
	p puintptr // attached p for executing go code (nil if not executing go code)

	//下一个要绑定到 M 的 P。当 M 执行完当前的 P 上的 Goroutine 后，会切换到 nextp 所指定的 P 上.
	nextp puintptr

	//用于保存在执行系统调用之前与 M（Machine）相关联的 P（Processor）。
	//当 M 需要执行系统调用时，它会从当前关联的 P 上分离，并保存原始的 P 地址到 oldp 字段。
	//这样，在系统调用完成后，M 可以重新关联到原始的 P 上，以继续执行 Go 代码。
	//puintptr 是一个 uintptr 类型的别名，用于保存指向 P 的地址。
	//通过使用 oldp 字段，M 可以在系统调用执行期间保存并恢复与 P 的关联关系。
	oldp puintptr // the p that was attached before executing a syscall

	//表示 M 的唯一标识符。每个 M 都有一个唯一的标识符，用于区分不同的 M。
	id int64

	//用于指示当前 M（Machine）是否正在执行 malloc 操作。
	//mallocing 字段的值为 1 表示当前 M 正在执行 malloc 操作，即分配内存。
	//如果 mallocing 字段的值为 0，则表示当前 M 不在执行 malloc 操作。
	//malloc 是动态内存分配的一种方式，在 Golang 中，它通常用于分配堆上的内存空间，用于存储动态创建的对象或数据结构。
	//通过使用 mallocing 字段，运行时系统可以跟踪每个 M 是否正在执行 malloc 操作。
	//这对于调度和资源管理非常重要，因为在执行 malloc 操作期间，运行时系统可能需要采取特殊的处理措施，
	//如处理内存分配失败、执行垃圾回收等。
	mallocing int32

	//用于指示当前 M（Machine）是否正在执行 panic 操作。
	//throwing 字段的值为 true 表示当前 M 正在执行 panic 操作，即触发异常。
	//通过使用 throwing 字段，运行时系统可以跟踪每个 M 是否正在执行 panic 操作。
	//这对于异常处理、调度和资源管理非常重要，因为在执行 panic 操作期间，运行时系统可能需要采取特殊的处理措施，
	//如处理异常、执行栈展开等。
	throwing throwType

	//用于指示是否禁用抢占（preemption）并保持当前 Goroutine 在该 M（Machine）上继续运行。
	//preemptoff 字段的值如果不为空字符串（""），则表示禁用抢占并保持当前 Goroutine 在该 M 上继续执行。
	//如果 preemptoff 字段的值为空字符串，则表示抢占机制处于启用状态。
	//抢占是指在多线程编程中，当一个线程正在执行时，操作系统或调度器中断该线程的执行，
	//并将 CPU 时间片分配给其他优先级更高的线程。
	//这是为了确保公平的线程调度和避免某个线程长时间占用 CPU 资源。
	//通过使用 preemptoff 字段，Golang 的运行时系统可以临时禁用抢占机制，并保持当前 Goroutine 在该 M 上继续执行。
	//这可以用于一些特殊的场景，例如在关键代码段或特定的调试操作中禁用抢占。
	preemptoff string // if != "", keep curg running on this m

	//用于表示当前 M（Machine）持有的锁的数量。
	//当 M 获取一个锁时，locks 字段的值会增加；当 M 释放一个锁时，locks 字段的值会减少。
	//通过使用 locks 字段，运行时系统可以跟踪每个 M 持有的锁的数量。这对于调度和资源管理非常重要，因为锁的获取和释放可能会影响 Goroutine 的调度和执行
	locks int32

	//用于表示当前 M（Machine）是否正在退出或终止的状态。
	//dying 字段的值为非零表示 M 正在退出或终止的过程中。
	//当 M 准备退出时，它会将 dying 字段的值设置为非零，并且在退出前进行必要的清理操作。
	//M 的退出通常发生在运行时系统关闭或 M 不再需要的情况下。在退出过程中，M 可能会释放资源、关闭连接、执行清理操作等。
	//通过使用 dying 字段，运行时系统可以跟踪每个 M 是否处于退出状态，并在必要时进行适当的清理和处理。
	//需要注意的是，dying 字段是在运行时系统内部使用的，对于用户代码来说是不可见的。它用于运行时系统的 M 管理和资源清理。
	dying int32

	//表示与 M 相关的分析器的采样频率
	profilehz int32

	//表示 M 当前是否处于自旋状态。
	//自旋是一种等待的方式，当 M 需要等待一些条件满足时，会快速地尝试获取锁或资源，而不进行阻塞。
	spinning bool // m is out of work and is actively looking for work

	//用于表示当前 M（Machine）是否被阻塞在一个 note 上。
	//note 是运行时系统中用于协调和同步 Goroutine 的一种基本机制。
	//当一个 Goroutine 需要等待某个事件或条件满足时，它会被阻塞在一个 note 上，直到该事件发生或条件满足。
	//blocked 字段的值为 true 表示当前 M 正在被阻塞在一个 note 上。
	//通过使用 blocked 字段，运行时系统可以跟踪每个 M 是否被阻塞，并相应地进行调度和资源管理。
	//当一个 M 被阻塞时，运行时系统可以将其从可运行的 M 集合中移除，并在事件发生后重新调度该 M
	blocked bool // m is blocked on a note

	//newSigstack用于表示是否在 C 线程上调用了 sigaltstack 函数执行 minit（M 的初始化）。
	//sigaltstack 是一个操作系统提供的函数，用于设置替代信号栈（alternate signal stack）。
	//在 Golang 的运行时系统中，M 的初始化过程（称为 minit）需要在一个特定的线程上执行，通常是 C 线程。
	//newSigstack 字段的值为 true 表示在执行 minit 时已在 C 线程上调用了 sigaltstack 函数，从而设置了替代信号栈。
	//如果 newSigstack 字段的值为 false，则表示没有调用 sigaltstack 函数。
	//通过使用 newSigstack 字段，运行时系统可以知道是否已设置了替代信号栈，并相应地调整 M 的初始化过程。
	//替代信号栈的设置可以确保在处理异常或信号时，运行时系统能够在安全的上下文中执行相应的处理。
	newSigstack bool // minit on C thread called sigaltstack

	//用于表示当前 M（Machine）是否持有打印锁（print lock）。 //打印锁是用于保护打印操作的一种锁机制。
	//在 Golang 的运行时系统中，当多个 Goroutine 同时进行打印操作时，需要使用打印锁来确保打印操作的原子性和顺序性。
	//printlock 字段的值为非零表示当前 M 持有打印锁。如果 printlock 字段的值为零，则表示当前 M 不持有打印锁。
	//通过使用 printlock 字段，运行时系统可以在进行打印操作时进行同步和互斥。
	//当一个 M 持有打印锁时，其他 M 需要等待打印锁被释放才能进行打印操作。
	printlock int8

	//用于表示当前 M（Machine）是否正在执行一个 cgo 调用。
	//通过使用 incgo 字段，运行时系统可以跟踪每个 M 是否正在执行 cgo 调用。
	//这对于调度和资源管理非常重要，因为 cgo 调用可能涉及与 C 代码的交互、资源管理的不同策略以及调度的特殊需求。
	//通过使用 incgo 字段，Golang 的运行时系统可以跟踪每个 M 是否正在执行 cgo 调用，
	//并在需要时进行适当的调度和资源管理，以确保与 C 代码的交互正确进行和资源的正确释放。
	incgo bool // m is executing a cgo call

	//isextra用于表示当前 M（Machine）是否是额外的（extra）M。
	//在 Golang 的并发模型中，M 是用于执行 Goroutine 的调度和管理的单位。
	//通常情况下，每个逻辑处理器（P）关联一个 M。
	//然而，当系统负载较低或需要更多并发性时，额外的 M 可以被创建来充分利用可用的处理资源。
	//isextra 字段的值为 true 表示当前 M 是一个额外的 M。反之，则表示当前 M 不是额外的 M。
	//通过使用 isextra 字段，运行时系统可以识别哪些 M 是额外的，并相应地进行调度和资源管理。
	//额外的 M 可以在需要更多并发性时被激活，而在负载较低时可以被休眠或销毁，以优化系统的性能和资源利用率。
	isextra bool // m is an extra m

	//用于表示是否可以安全地释放 g0 和删除 M（Machine）。
	//在 Golang 的运行时系统中，g0 是一个特殊的 Goroutine，它用于执行系统级任务和初始化操作。
	//M 则是与 Goroutine 相关联的执行上下文。当 M 不再需要时，可以尝试释放与之关联的资源。
	//freeWait 字段是一个原子类型的整数，用于记录释放 M 的安全状态。
	//当 freeWait 字段的值为非零时，表示不安全释放 M；当 freeWait 字段的值为零时，表示可以安全释放 M。
	//通过使用 freeWait 字段，运行时系统可以跟踪 M 是否可以安全地释放，并在需要时进行适当的处理。
	//这通常涉及等待相关的操作完成，以确保在释放 M 之前不会出现竞争条件或数据访问冲突。
	freeWait atomic.Uint32 // Whether it is safe to free g0 and delete m (one of freeMRef, freeMStack, freeMWait)

	//用于表示快速随机数生成器的状态。
	//fastrand 字段用于在运行时系统中生成快速的伪随机数。它是一个无锁的随机数生成器，用于产生高性能的伪随机数序列。
	//通过使用 fastrand 字段，运行时系统可以在需要快速随机数生成的场景中，
	//如哈希函数、并发算法、随机算法等，快速获取伪随机数。
	fastrand uint64

	//用于表示是否需要额外的 M（Machine）。 needextram 字段的值为 true 表示当前 M 需要额外的 M
	needextram bool

	//用于表示是否需要进行 Goroutine 的回溯（traceback）
	//traceback 字段的值表示回溯的级别或模式。它用于控制在发生错误或异常时，运行时系统是否要生成 Goroutine 的调用栈跟踪信息。
	//通过使用 traceback 字段，Golang 的运行时系统可以根据指定的级别生成适当的回溯信息，以帮助开发人员定位问题并进行调试。
	//这提供了一种方便的方法来捕获和分析发生在 Goroutine 中的错误和异常。
	//通过使用 traceback 字段，可以指定需要生成的回溯级别，如完整回溯、简化回溯或禁用回溯等。
	//这有助于在调试和错误排查时获取有用的调用栈信息。
	traceback uint8

	//用于表示发生的 cgo 调用的总次数。
	//通过使用 ncgocall 字段，可以统计和监控程序中的 cgo 调用次数，以便在性能优化、调试和资源管理等方面进行分析和优化。
	ncgocall uint64 // number of cgo calls in total

	//用于表示当前正在进行中的 cgo 调用的数量。
	//cgo 调用是指在 Golang 代码中调用 C 语言函数或在 C 语言代码中调用 Golang 函数的操作。
	//在进行 cgo 调用时，涉及到在 Golang 和 C 语言之间的上下文切换和数据传递。
	//ncgo 字段用于跟踪当前正在进行中的 cgo 调用的数量。它反映了程序中正在执行的与 C 语言相关的操作的规模和并发性。
	//通过使用 ncgo 字段，可以了解当前有多少个 cgo 调用正在同时进行，这对于了解程序中的并发性和资源利用情况是有帮助的。
	//需要注意的是，ncgo 字段是在运行时系统内部使用的，对于用户代码来说是不可见的。它用于运行时系统的性能分析和资源管理。
	//通过使用 ncgo 字段，Golang 的运行时系统可以提供有关当前正在进行中的 cgo 调用数量的信息，
	//以帮助开发人员识别并发性问题、调整并发级别以及进行资源分配和调度。
	ncgo int32 // number of cgo calls currently in progress

	//用于表示 cgoCallers 是否被临时使用
	//cgoCallers 是用于跟踪 cgo 调用的调用栈信息的数据结构。它用于记录 cgo 调用的调用栈，以便在需要时进行调试和跟踪。
	//cgoCallersUse 字段的值为非零表示 cgoCallers 正在被临时使用。
	//通过使用 cgoCallersUse 字段，运行时系统可以在 cgo 调用期间临时使用 cgoCallers 数据结构，
	//并确保在使用过程中的正确性和线程安全性。
	cgoCallersUse atomic.Uint32 // if non-zero, cgoCallers in use temporarily

	//cgoCallers 用于记录在 cgo 调用过程中发生崩溃时的回溯信息.
	//cgoCallers 结构体用于跟踪 cgo 调用期间的调用栈，以便在发生崩溃或异常时提供关于崩溃点的调用栈信息。
	//通过使用 cgoCallers 字段，运行时系统可以将崩溃点与 cgo 调用相关联，并在需要时提供相关的调用栈信息，
	//以帮助开发人员定位和解决崩溃问题。
	cgoCallers *cgoCallers // cgo traceback if crashing in cgo call

	//park用于实现 Goroutine 的休眠（park）和唤醒操作。它是note类型。
	//note 是运行时系统中用于协调和同步 Goroutine 的基本机制之一。
	//当一个 Goroutine 需要等待某个事件或条件满足时，它会被休眠在一个 note 上，直到被其他 Goroutine 唤醒。
	//park 字段表示当前 M（Machine）所关联的 Goroutine 是否被休眠。
	//如果 park 字段的值为非零，则表示当前 Goroutine 处于休眠状态；如果 park 字段的值为零，则表示当前 Goroutine 是可运行的。
	//通过使用 park 字段，运行时系统可以对 Goroutine 进行休眠和唤醒的操作。
	//当 Goroutine 需要等待某个事件时，它会被休眠在相关的 note 上，直到事件发生。
	//当事件发生时，其他 Goroutine 可以通过操作相关的 note 来唤醒休眠的 Goroutine。
	park note

	//alllink 字段是一个指向 m 结构体的指针，用于维护在 allm 列表上的所有 M（Machine）的链接。
	//allm 列表是运行时系统中的一个全局列表，用于存储所有的 M。
	//每个 M 在启动时会被添加到 allm 列表中，以便运行时系统能够轻松地遍历和操作所有的 M
	alllink *m // on allm

	//schedlink 字段是一个 muintptr 类型的字段，用于在调度器的链表中链接 M。调度器链表维护了一组可运行的 M，它们处于待调度状态。当一个 M 可以被调度执行时，它将链接到调度器链表中，以便调度器可以轮询和选择要执行的 M。schedlink 字段的作用是将当前 M 链接到调度器链表中。
	schedlink muintptr

	//表示在执行系统调用或其他阻塞操作期间，M 持有的锁相关的 Goroutine。
	//当 M 需要阻塞时，会将此字段设置为需要持有的锁所在的 Goroutine。
	lockedg guintptr

	//用于记录创建当前线程的堆栈信息.
	//通过使用 createstack 字段，运行时系统可以追踪创建当前 M 的线程的堆栈信息，以便在需要时进行调试和分析。
	//这对于了解线程的创建和执行环境是很有帮助的。
	createstack [32]uintptr // stack that created this thread.

	//用于跟踪外部 LockOSThread 操作的状态。
	//lockOSThread 是一个用于将当前线程锁定到当前的 M（Machine）的操作。
	//当调用 lockOSThread 后，该线程将与当前的 M 关联，即只有该线程才能执行与该 M 相关的 Goroutine。
	//lockedExt 字段用于跟踪外部 LockOSThread 操作的状态。当 lockedExt 的值为非零时，表示当前 M 的线程已被外部锁定；
	//当 lockedExt 的值为零时，表示当前 M 的线程未被外部锁定。
	//通过使用 lockedExt 字段，运行时系统可以追踪外部 LockOSThread 操作的状态，并确保在适当的时候将外部线程与 M 关联。
	lockedExt uint32 // tracking for external LockOSThread

	//用于跟踪内部 lockOSThread 操作的状态。

	//lockedInt 字段用于跟踪 lockOSThread 操作的状态。
	//当 lockedInt 的值为非零时，表示当前 M 的线程已被锁定到该 M；当 lockedInt 的值为零时，表示当前 M 的线程未被锁定。
	//通过使用 lockedInt 字段，运行时系统可以追踪 lockOSThread 操作的状态，并确保在适当的时候将线程与 M 关联。
	lockedInt uint32 // tracking for internal lockOSThread

	//用于表示下一个正在等待锁的 M（Machine）。
	//当一个 M 试图获取一个已被其他 M 锁定的资源时，它将被阻塞并加入到等待队列中。
	//nextwaitm 字段用于维护等待队列中下一个等待锁的 M 的引用。
	//通过使用 nextwaitm 字段，运行时系统可以维护 M 等待锁的顺序，并确保正确的调度和资源分配。
	nextwaitm muintptr // next m waiting for lock

	//waitunlockf 字段指向一个函数，该函数用于在 Goroutine 等待解锁时执行特定的操作。
	//通过使用 waitunlockf 字段，运行时系统可以在 Goroutine 等待解锁时执行自定义的操作。
	//这可能涉及到 Goroutine 的状态变更、唤醒等相关操作
	waitunlockf func(*g, unsafe.Pointer) bool

	//waitlock 字段用于在 Goroutine 等待解锁时存储与该等待操作相关的信息。
	//它可以是一个锁对象、条件变量或其他用于同步的数据结构。
	//通过使用 waitlock 字段，运行时系统可以在 Goroutine 等待解锁时将相关的等待状态和条件存储在 waitlock 字段中，
	//以便在解锁时能够恢复相关的等待操作。
	waitlock unsafe.Pointer

	//用于表示在等待状态下的追踪事件类型。
	//它可以用于跟踪和记录在 Goroutine 等待期间发生的事件，例如等待锁的持续时间、唤醒时机等
	waittraceev byte

	//waittraceskip 字段用于指定在追踪 Goroutine 等待状态时要跳过的堆栈帧数。
	//它可以用于在追踪等待状态时忽略某些不需要的堆栈帧，以减少追踪的开销。
	//通过使用 waittraceskip 字段，运行时系统可以根据需要设置要跳过的堆栈帧数，以适应不同的调试和分析需求。
	waittraceskip int

	//用于指示当前 M（Machine）是否处于启动追踪状态。
	//startingtrace 字段用于跟踪 M 的启动过程中是否正在进行追踪。当 startingtrace 的值为 true 时，表示当前 M 正在进行启动追踪；
	//通过使用 startingtrace 字段，运行时系统可以在 M 启动过程中追踪特定的事件或操作，以帮助调试和分析启动过程中的问题
	startingtrace bool

	//用于记录系统调用的计时器。syscalltick 字段用于跟踪 M（Machine）执行系统调用的次数。
	//每当 M 执行一个系统调用时，syscalltick 字段的值会递增。
	//通过使用 syscalltick 字段，运行时系统可以统计 M 执行系统调用的频率，并用于调度和性能分析。
	syscalltick uint32

	//用于将当前的 M（Machine）链接到 sched.freem 列表上。
	//sched.freem 列表是调度器中的一个空闲 M 列表，用于存储处于空闲状态的 M。
	//当一个 M 不再被使用时，它会被链接到 sched.freem 列表上，以便稍后可以被重新分配给其他的 Goroutine。
	//通过使用 freelink 字段，运行时系统可以将当前的 M 链接到 sched.freem 列表上，以便将其标记为可用的空闲 M。
	freelink *m // on sched.freem

	//libcall用于在底层的 NOSPLIT 函数中存储太大而无法放置在栈上的数据，
	//以避免将其放置在栈上造成栈溢出或栈帧过大的问题。
	//通过使用 libcall 字段，运行时系统可以在需要的时候为底层的 NOSPLIT 函数分配额外的内存空间，以容纳较大的数据。
	libcall libcall

	//libcallpc 字段用于记录底层 NOSPLIT 函数中的 libcall 调用的程序计数器（PC）值。
	//CPU 分析器使用这个字段来跟踪底层函数的执行情况，以便分析函数的性能和运行时间。
	//通过使用 libcallpc 字段，运行时系统可以在底层 NOSPLIT 函数中跟踪和记录 CPU 分析信息，
	//从而提供有关函数执行时间和性能的详细数据。
	libcallpc uintptr // for cpu profiler

	//libcallsp 字段用于记录底层 NOSPLIT 函数中的 libcall 调用的栈指针的位置。
	//它可以帮助运行时系统正确地管理和恢复栈状态，以确保底层函数的正常执行。
	//通过使用 libcallsp 字段，运行时系统可以在底层 NOSPLIT 函数中跟踪和管理栈的状态，以支持 libcall 的调用和返回。
	libcallsp uintptr

	//libcallg 字段用于记录底层 NOSPLIT 函数中的 libcall 调用所关联的 Goroutine。
	//它可以帮助运行时系统跟踪和管理底层函数中涉及的 Goroutine，以确保正确的并发调度和协作。
	//通过使用 libcallg 字段，运行时系统可以在底层 NOSPLIT 函数中关联正确的 Goroutine，
	//以便正确处理 Goroutine 的上下文和状态。
	libcallg guintptr

	//用于在 Windows 系统上存储系统调用的参数。
	//在 Windows 系统上，进行系统调用时需要将参数保存在 syscall 字段中。
	syscall libcall // stores syscall parameters on windows

	//用于在 VDSO 调用期间进行回溯时的栈指针（SP）位置。
	//VDSO（Virtual Dynamic Shared Object）是一个虚拟共享对象，它由操作系统内核提供，
	//用于提供一些常见的系统调用功能，以减少用户态和内核态之间的切换次数。
	//在某些情况下，Golang 的运行时系统可以利用 VDSO 来加速一些系统调用的执行。
	//vdsoSP 字段用于记录当前 Goroutine 在进行 VDSO 调用时的栈指针位置。
	//它可以帮助运行时系统在进行 VDSO 调用期间进行正确的回溯，并获取与 VDSO 调用相关的堆栈跟踪信息
	vdsoSP uintptr // SP for traceback while in VDSO call (0 if not in call)

	//用于在 VDSO 调用期间进行回溯时的程序计数器（PC）位置。
	//vdsoPC 字段用于记录当前 Goroutine 在进行 VDSO 调用时的程序计数器位置。
	//跟vdsoSP一样，它可以帮助运行时系统在进行 VDSO 调用期间进行正确的回溯，并获取与 VDSO 调用相关的堆栈跟踪信息。
	vdsoPC uintptr // PC for traceback while in VDSO call

	//  preemptGen 字段是一个 atomic.Uint32 类型的原子变量，用于计数已完成的抢占信号。
	//preemptGen 字段用于记录已完成的抢占信号的数量。
	//抢占信号是用于触发 Goroutine 抢占的机制，当一个 Goroutine 被抢占时，会增加 preemptGen 字段的计数。
	//通过使用 preemptGen 字段，运行时系统可以检测到抢占请求是否成功。
	//如果抢占请求失败，即 preemptGen 字段的计数没有增加，可以得知抢占操作未能生效。
	preemptGen atomic.Uint32

	// 用于表示当前 M（Machine）上是否存在挂起的抢占信号.
	//当一个 Goroutine 请求抢占时，抢占信号会被设置为挂起状态，表示该 M 正在等待抢占。
	//通过使用 signalPending 字段，运行时系统可以检查当前 M 是否存在挂起的抢占信号，以便在合适的时机触发抢占操作。
	//需要注意的是，signalPending 字段是一个原子变量，用于在并发环境下进行安全的状态标记，以避免竞态条件。
	//通过使用 signalPending 字段，Golang 的运行时系统可以准确地检测和处理抢占信号的挂起状态，以支持可靠的并发调度和资源管理。
	signalPending atomic.Uint32

	//dlogPerM 是一个布尔类型的常量，用于指示是否为每个 M（Machine）启用调试日志（debug log）。
	//当 dlogPerM 为 true 时，每个 M 都会有自己独立的调试日志。这意味着每个 M 都会维护和输出其自己的调试日志信息。
	//当 dlogPerM 为 false 时，所有的 M 共享同一个调试日志。这意味着所有 M 的调试日志信息会被写入到同一个日志中。
	//需要注意的是，dlogPerM 是一个常量，用于在编译时确定调试日志的配置方式。
	//它的设置通常是在运行时系统的构建配置中进行定义。
	//通过使用 dlogPerM 常量，Golang 的运行时系统可以根据需要配置和管理调试日志，以帮助开发人员调试和排查问题。
	dlogPerM

	//表示操作系统相关的信息和状态。不同操作系统所存储的信息会有差异
	mOS

	//locksHeldLen 字段用于记录当前 M 持有的锁的数量，最多可以持有 10 个锁。
	//该字段由锁排序代码（lock ranking code）维护和更新。
	//通过使用 locksHeldLen 字段，运行时系统可以追踪和管理当前 M 持有的锁的数量，
	//以便在锁的获取和释放过程中进行适当的调度和优化。
	locksHeldLen int

	// 用于存储当前 M（Machine）持有的锁的信息。
	//locksHeld 字段是一个长度为 10 的数组，每个元素都是 heldLockInfo 类型的结构体，
	//用于记录锁的详细信息，例如锁的地址、持有者 Goroutine 的信息等。
	//通过使用 locksHeld 字段，运行时系统可以跟踪和管理当前 M 持有的锁的信息，以便进行锁的获取、释放和调度。
	//锁的信息可以帮助运行时系统优化并发调度策略，避免死锁和竞态条件。
	locksHeld [10]heldLockInfo
}
