package gmp

import (
	"sync/atomic"
)

// 取 processor 的首字母，为 M 的执行提供“上下文”，保存 M 执行 G 时的一些资源，
// 例如本地可运行 G 队列，memeory cache 等。
// 一个 M 只有绑定 P 才能执行 goroutine，
// 当 M 被阻塞时，整个 P 会被传递给其他 M ，或者说整个 P 被接管
type p struct {
	// 在 allp 中的索引,通过 id 字段，可以在 allp 数组中快速定位到对应的 p 结构体，以进行处理器级别的操作和管理。
	//allp 是一个全局的 p 数组，存储了系统中所有的处理器（Processor）的信息。
	id int32

	//用于表示处理器（Processor）的状态。状态值有_Pidle、_Prunning、_Psyscall、_Pgcstop 和 _Pdead
	//_Pidle = 0	处理器没有运行用户代码或者调度器，被空闲队列或者改变其状态的结构持有，运行队列为空；也有可能是几种状态正在过度中状态
	//_Prunning = 1	被线程 M 持有，并且正在执行用户代码或者调度器。只能由拥有当前P的M才可能修改此状态。M可以将P的状态修改为_Pidle（无工作可做）、_Psyscall(系统调用) 或 _Pgstop(GC); 另外M也可以P的使用权交给另一个M（调度一个锁定状态的G）
	//_Psyscall = 2	当前P没有执行用户代码，当前线程陷入系统调用
	//_Pgcstop =3	被线程 M 持有，当前处理器由于垃圾回收被停止，由 _Prunning 变为 _Pgcstop
	//_Pdead = 4	 当前处理器已经不被使用，如通过动态调小 GOMAXPROCS 进行 P 收缩
	status uint32 // one of pidle/prunning/...

	//用于将 P（Processor）组织成链表。
	//puintptr 实际上是一个 uintptr 类型的别名，它表示一个指向另一个 P 的指针。
	//通过将多个 P 的 link 字段连接在一起，可以形成一个 P 链表。这种链表通常被用于调度和管理可用的处理器资源。
	//在调度器中，当一个 Goroutine 需要被分配到一个 P 上执行时，调度器会遍历 P 链表，查找可用的 P，并将 Goroutine 分配给该 P。
	//link 字段使得 P 可以在链表中进行连接和断开，以便于调度器根据需要动态地分配和回收处理器资源。
	link puintptr

	// 每次调用 schedule 时会加一.用于调度器的统计和性能分析.
	//在 Golang 的调度器中，调度器会周期性地检查各个 P（Processor）的队列，并选择可执行的 Goroutine 进行调度。
	//每次进行调度时，都会增加 schedtick 字段的值，以记录调度操作的次数。
	schedtick uint32 // incremented on every scheduler call

	// 每次系统调用时加一.
	//Golang 的运行时系统可以跟踪和统计系统调用的次数，从而评估系统调用的开销和性能状况，并进行相应的优化和调整。
	syscalltick uint32 // incremented on every system call

	// 用于跟踪监视器观察到的时钟滴答。时钟滴答是一个在调度器中使用的时间单位，用于衡量 Goroutine 的执行时间和调度情况。
	//监视器是 Golang 运行时系统中的一个组件，用于监控 Goroutine 的行为和性能，并进行相应的调整和优化。
	//通过监视器，运行时系统可以检测出长时间运行的 Goroutine、阻塞的 Goroutine 等情况，并采取适当的措施来平衡调度和资源利用。
	sysmontick sysmontick // last tick observed by sysmon

	// 用于指向与当前 P（Processor）关联的 M（Machine）.
	//m 字段用于建立 P 和 M 之间的关联关系。每个 P 都会关联一个 M，表示该 P 当前正在执行的 M。
	//当 P 需要执行 Goroutine 时，它会从与之关联的 M 中获取一个可执行的 Goroutine，并将其分配给自己执行。
	//需要注意的是，如果 P 处于空闲状态（_Pidle），即没有可执行的 Goroutine，那么 m 字段将为空指针（nil）。
	//通过使用 m 字段，P 和 M 实现了双向关联，使得运行时系统能够将 Goroutine 均匀地分配给各个处理器，以实现并行的执行。
	m muintptr // back-link to associated m (nil if idle)

	//mcache 是每个 M（Machine）与本地缓存相关的数据结构。它包含了与 M 关联的 Goroutine 执行时所需的本地分配缓存。
	//每个 M 都有自己的本地缓存（mcache）来加速 Goroutine 的内存分配和回收。
	//当 Goroutine 需要分配内存时，会首先检查关联的 M 的 mcache 是否为空。
	//如果不为空，则直接从 mcache 中获取内存进行分配。这样可以避免频繁地向全局堆申请内存，提高分配速度。
	//mcache 中还包含了分配和回收内存的一些其他信息，例如空闲对象列表、本地缓存的大小等。
	mcache *mcache

	//pcache 字段是一个 pageCache 类型的变量。它是一种无锁的缓存，用于分配内存页。
	//页面缓存的作用是提供一块内存区域，供分配器在没有锁的情况下分配内存页。
	//页面缓存是针对每个 P 的，每个 P 都有自己的页面缓存。
	//当分配器需要分配内存页时，会首先查找页面缓存，检查其中是否有空闲的页面可用。
	//如果有空闲页面，则直接分配给请求者，并更新位图表示页面状态。如果没有空闲页面，则需要从全局堆申请新的内存页。
	//通过使用页面缓存，可以避免频繁地向操作系统申请内存页，提高内存分配和释放的效率。
	//虽然 pcache 和 mcache 都与内存管理相关，但它们服务于不同的目的和使用场景。
	//pcache 主要用于缓存页面，以减少频繁向操作系统申请内存页的开销，提高内存分配和释放的效率。
	//它与 P 相关联，用于 P 的内存管理。
	//mcache 则主要用于本地分配缓存，以加速 Goroutine 的内存分配和回收操作。它与 M 相关联，用于 M 的内存管理。
	//在运行时系统中，P 和 M 是并发执行的基本单位，它们负责不同的任务和功能。
	//pcache 和 mcache 是为了支持它们的不同需求而设计的
	pcache pageCache

	//raceprocctx 字段用于与竞态检测（race detection）相关的上下文信息。竞态检测是一种用于检测并发程序中数据竞争问题的工具。
	//在启用竞态检测时，每个 Golang 程序都会关联一个竞态检测上下文，用于跟踪和记录竞态检测的相关信息。
	//raceprocctx 字段存储了与竞态检测相关的上下文信息，通常是一个指向特定上下文数据结构的指针。
	//竞态检测工具会在程序运行过程中检测并发访问共享数据的情况，通过比较访问的顺序和时间戳等信息来发现可能存在的数据竞争。
	//当检测到数据竞争时，竞态检测器会使用竞态检测上下文来记录相关的信息，以便进行进一步的分析和报告。
	raceprocctx uintptr

	//deferpool 是一个 _defer 指针的切片，用于存储可用的延迟调用结构体
	//deferpool 和 deferpoolbuf 用于管理 _defer 结构体的缓存池。
	//deferpool也是一种池化技术，缓存池是为了避免频繁地分配和释放 _defer 结构体而设计的，以提高性能和效率。
	//当函数执行过程中遇到延迟调用时，会从 deferpool 中获取一个可用的 _defer 结构体。
	//如果 deferpool 中没有可用的结构体，则会动态分配一个新的 _defer 结构体。
	//执行完延迟调用后，将 _defer 结构体放回 deferpool 中，以便重复使用。
	//通过使用缓存池，可以减少动态内存分配的次数，从而提高程序的性能和效率。
	deferpool []*_defer // pool of available defer structs (see panic.go)

	//deferpoolbuf 是一个长度为 32 的 _defer 指针数组.
	//当函数执行过程中遇到延迟调用时，会先尝试从 deferpoolbuf 中获取一个可用的 _defer 结构体。
	//如果 deferpoolbuf 中有可用的结构体，则会使用该结构体来存储延迟调用的信息。
	//如果 deferpoolbuf 中没有可用的结构体，则会动态分配一个新的 _defer 结构体来存储延迟调用的信息。
	//通过使用 deferpoolbuf，可以减少动态内存分配的次数，提高延迟调用的效率和性能。
	//deferpoolbuf与deferpool的区别？
	//deferpoolbuf 是一个静态的数组，用于存储 _defer 结构体的缓冲区。
	//它的大小是固定的，并且通常比较小，因此适用于频繁创建和释放延迟调用的场景。
	//deferpool 是一个动态的切片，用于存储可用的 _defer 结构体。它的大小会根据需要动态调整，以适应不同的延迟调用的数量。
	//它主要用于在运行时系统中管理和重用 _defer 结构体，以减少内存分配的开销。、
	//即deferpoolbuf用于频繁创建和释放延迟调用的场景，而deferpool在于重用_defer
	deferpoolbuf [32]*_defer

	// goidcache 和 goidcacheend 是用于缓存 Goroutine ID 的字段。
	//为了提高获取 Goroutine ID 的性能，runtime 包中维护了一个 Goroutine ID 的缓存。
	//这个缓存存储在 goidcache 和 goidcacheend 字段中。
	//goidcache 是一个 uint64 类型的字段，表示 Goroutine ID 缓存的起始值。
	//当程序需要获取 Goroutine ID 时，首先检查缓存中是否有可用的值。 如果有，就直接使用缓存中的值
	//goidcacheend 是一个 uint64 类型的字段，表示 Goroutine ID 缓存的结束值。
	//当缓存中的 Goroutine ID 达到 goidcacheend 时，需要重新获取新的 Goroutine ID 并更新缓存。
	//通过使用 Goroutine ID 缓存，可以减少对 runtime·sched.goidgen 的访问次数，从而提高获取 Goroutine ID 的效率。
	goidcache    uint64
	goidcacheend uint64

	// runqhead 和 runqtail 是用于表示可运行 Goroutine 队列的字段。
	//可运行 Goroutine 队列用于存储可立即执行的 Goroutine，即那些已经准备好被调度执行的 Goroutine。
	//runqhead 是一个 uint32 类型的字段，表示可运行 Goroutine 队列的头部位置。
	//当需要从可运行队列中获取 Goroutine 进行调度时，会从 runqhead 指定的位置开始获取。
	//runqtail 是一个 uint32 类型的字段，表示可运行 Goroutine 队列的尾部位置。
	//当有新的 Goroutine 准备好被调度时，会将其添加到 runqtail 指定的位置。
	//通过使用 runqhead 和 runqtail，运行时系统可以快速访问可运行 Goroutine 队列，实现高效的 Goroutine 调度。
	runqhead uint32
	runqtail uint32

	//runq 数组用于存储可运行 Goroutine 的指针。
	//guintptr 表示一个 g（Goroutine）的指针。当一个 Goroutine 准备好被调度时，它的指针会被添加到 runq 队列中。
	//每个P都有一个自己的runq，除了自身有的runq 还有一个全局的runq, 对于每个了解过GPM的gohper应该都知道这一点。
	//每个P下面runq的允许的最大goroutine 数量为256。
	runq [256]guintptr

	// runnext 非空时，代表的是一个 可运行状态的 G，
	// 这个 G 被 当前 G 修改为 ready 状态，相比 runq 中的 G 有更高的优先级。
	// 如果当前 G 还有剩余的可用时间，那么就应该运行这个 G
	// 运行之后，该 G 会继承当前 G 的剩余时间. 这个字段是用来实现调度器亲和性的，
	//当一个G阻塞时，这时P会再获取一个G进行绑定执行，
	//如果这时原来的G执行阻塞结束后，如果想再次被接着继续执行，就需要重新在P的 runq 进行排队，
	//当 runq 里有太多的goroutine 时将会导致这个刚刚被解除阻塞的G迟迟无法得到执行， 同时还有可能被其他处理器所窃取。
	//从 Go 1.5 开始得益于 P 的特殊属性，从阻塞 channel 返回的 Goroutine 会优先运行，
	//这里只需要将这个G放在 runnext 这个字段即可。
	runnext guintptr

	// gFree 是一个用于管理空闲 Goroutine 的数据结构。
	//gFree 结构体包含以下字段：
	//gList：一个双向链表，用于存储空闲的 Goroutine。每个空闲 Goroutine 在链表中都有一个节点，可以通过该节点访问链表中的其他空闲 Goroutine。
	//n：一个 int32 类型的字段，表示空闲 Goroutine 的数量。
	//当一个 Goroutine完成执行（状态为 Gdead）并且不再需要继续使用时，它会被添加到 gFree 的链表中，以便后续可以被重新利用。
	//通过维护一个空闲 Goroutine 的链表，运行时系统可以节省创建和销毁 Goroutine 的开销，并且能够快速地获取可用的 Goroutine。
	//当需要创建新的 Goroutine时，运行时系统会首先尝试从 gFree 链表中获取一个空闲 Goroutine。
	//如果链表为空，那么会动态分配一个新的 Goroutine。
	//当一个 Goroutine完成执行后，它会被释放并添加到 gFree 链表中，以便下次需要时可以重复使用。
	//通过使用 gFree 数据结构，Golang 的运行时系统可以高效地管理和重用空闲的 Goroutine，
	//从而提高程序的并发性能和资源利用率。
	gFree struct {
		gList
		n int32
	}

	//sudogcache 和 sudogbuf 是用于管理和缓存 sudog 结构体的字段。
	//sudog 是用于表示等待队列中的 Goroutine 的数据结构。
	//sudogcache 是一个 []*sudog 切片，用于缓存可用的 sudog 结构体。
	//这个缓存池用于避免频繁地分配和释放 sudog 结构体，以提高性能和效率。
	//当一个 sudog 结构体不再使用时，会将其放回 sudogcache 中，以便下次需要时可以重复使用。
	//sudogbuf 是一个长度为 128 的 *sudog 数组，用于存储 sudog 结构体的缓冲区。
	//当需要创建新的 sudog 结构体时，首先会尝试从 sudogcache 中获取一个可用的结构体。
	//如果缓存中有可用的结构体，则会使用它来存储 sudog 的信息。
	//如果缓存中没有可用的结构体，则会动态分配一个新的 sudog 结构体。
	//通过使用 sudogcache 和 sudogbuf，可以减少动态内存分配的次数，提高 sudog 结构体的创建和释放的效率。
	sudogcache []*sudog
	sudogbuf   [128]*sudog

	// mspancache 是用于缓存 mspan 对象的字段，它表示从堆中缓存的 mspan 对象。
	//mspan 是用于管理内存分配的数据结构，每个 mspan 对象代表了一块内存页的管理信息。
	//mspancache 结构体包含以下字段：
	//len：一个整型字段，表示当前缓存中 mspan 对象的数量。
	//buf：一个长度为 128 的 mspan 指针数组，用于存储 mspan 对象。
	//mspancache 用于缓存 mspan 对象，以提高内存分配的性能。
	//当需要分配新的内存页时，运行时系统首先会尝试从 mspancache 中获取一个可用的 mspan 对象。
	//如果缓存中有可用的对象，就会将其分配给新的内存页。这样可以减少对堆的访问，提高内存分配的效率。
	//需要注意的是，mspancache 的长度（len 字段）是显式声明的，
	//因为该字段在一些分配代码路径中可能涉及到不能使用写屏障的情况。因此，需要通过显式管理长度来避免写屏障的使用。
	//通过使用 mspancache，Golang 的运行时系统可以高效地管理和重用 mspan 对象，以提高内存分配和管理的性能。
	mspancache struct {
		// 一个整型字段，表示当前缓存中 mspan 对象的数量。
		len int
		//一个长度为 128 的 mspan 指针数组，用于存储 mspan 对象。
		buf [128]*mspan
	}

	//用于存储追踪信息的缓冲区
	tracebuf traceBufPtr

	//  用于指示是否需要跟踪垃圾回收器的扫描事件。
	//当 traceSweep 为 true 时，垃圾回收器会记录和追踪扫描事件。
	//这样可以提供详细的垃圾回收器执行信息，用于性能分析和调试。
	//具体来说，当需要进行垃圾回收时，如果 traceSweep 为 true，则会延迟触发扫描开始事件，
	//直到至少一个内存范围（span）被完全扫描完成。
	//这样可以确保垃圾回收器的扫描事件只包括实际进行扫描的内存范围。
	traceSweep bool
	// traceSwept 和 traceReclaimed 是用于跟踪垃圾回收器在当前扫描循环中扫描和回收的字节数的字段。
	//traceSwept 是一个 uintptr 类型的字段，用于跟踪当前扫描循环中已经扫描的字节数。它表示垃圾回收器已经检查并标记为可回收的内存字节数。
	//
	//traceReclaimed 是一个 uintptr 类型的字段，用于跟踪当前扫描循环中已经回收的字节数。
	//它表示垃圾回收器已经成功回收的内存字节数。
	//这两个字段的目的是提供有关垃圾回收器执行的详细信息，以便进行性能分析和调试。
	//通过跟踪已扫描和回收的字节数，可以评估垃圾回收器的效率和性能。
	traceSwept, traceReclaimed uintptr

	//palloc 是一个 persistentAlloc 类型的字段，用于每个 P（处理器）维护持久性分配（persistent allocation）以避免使用互斥锁。
	//persistentAlloc 是一种分配器，用于在特定的处理器上分配固定大小的内存块，而无需使用互斥锁进行同步。
	//它被设计为在每个 P 上独立运行，以避免并发访问的竞争条件。
	//通过在每个 P 上维护一个独立的 persistentAlloc 实例，Golang 运行时系统可以提高并发性能和分配的效率。
	//每个 P 都拥有自己的 persistentAlloc 实例，因此可以独立地进行分配操作，而无需在多个 P 之间进行同步。
	palloc persistentAlloc // per-P to avoid mutex

	// timer0When 是一个 atomic.Int64 类型的字段，用于表示定时器堆中第一个条目的触发时间。
	//定时器堆是一种数据结构，用于管理和调度定时器的触发时间。
	//每个定时器都会被插入到定时器堆中，并根据其触发时间进行排序。
	//timer0When 字段用于表示定时器堆中第一个条目的触发时间。
	//timer0When 字段是一个原子类型的整数，用于确保并发访问时的安全性。
	//它存储了定时器堆中第一个条目的触发时间，以纳秒为单位表示。如果定时器堆为空，timer0When 字段的值将为 0。
	//通过使用原子操作保护 timer0When 字段，运行时系统可以确保多个线程或 Goroutine 并发地访问和更新定时器堆的触发时间，
	//而不会发生竞争条件或数据不一致的情况。
	timer0When atomic.Int64

	// timerModifiedEarliest 是一个 atomic.Int64 类型的字段，
	//用于表示具有 "timerModifiedEarlier" 状态的定时器中最早的下一个触发时间。
	timerModifiedEarliest atomic.Int64

	// 用于表示在辅助分配（assistAlloc）中花费的纳秒数。辅助分配是垃圾回收期间由 P 执行的帮助分配操作。
	gcAssistTime int64 // Nanoseconds in assistAlloc

	//用于在分数化标记工作器（fractional mark worker）中花费的纳秒数。分数化标记工作器是垃圾回收期间执行部分标记的工作器。
	gcFractionalMarkTime int64 // Nanoseconds in fractional mark worker (atomic)

	//用于跟踪垃圾回收器 CPU 限制器的事件的字段
	// 在 Golang 的运行时系统中，垃圾回收器 CPU 限制器用于限制垃圾回收期间消耗的 CPU 时间，
	//以避免垃圾回收对应用程序的执行造成过大的影响。
	//limiterEvent 包含一个 stamp 字段，它是一个 atomic.Uint64 类型的变量。
	//stamp 字段用于存储一个 limiterEventStamp 值，该值表示事件的时间戳。
	//通过使用 limiterEvent，运行时系统可以跟踪和记录特定事件在时间上的发生情况。
	//这对于分析和优化垃圾回收过程中的 CPU 时间消耗非常有用。
	//通过收集事件的时间戳，可以计算出垃圾回收期间 CPU 时间的分布和利用率，从而帮助开发人员进行性能分析和调优
	limiterEvent limiterEvent

	// 表示下一个标记工作器（mark worker）运行的模式。
	//它用于与通过 gcController.findRunnableGCWorker 选择的工作器 Goroutine 进行通信。
	//在调度其他 Goroutine 时，必须将此字段设置为 gcMarkWorkerNotWorker。
	gcMarkWorkerMode gcMarkWorkerMode

	// 最近一个标记工作器开始运行的时间，以纳秒为单位。
	gcMarkWorkerStartTime int64

	// gcw is this P's GC work buffer cache. The work buffer is
	// filled by write barriers, drained by mutator assists, and
	// disposed on certain GC state transitions.
	//  P 的 GC工作的缓冲区缓存。
	// GC工作的缓冲区由写屏障填充、由助理分配器（mutator assists）消耗，并在某些 GC 状态转换时释放。
	gcw gcWork

	//  P 的 GC 写屏障缓存，后续版本可能会考虑缓存正在运行的G。
	wbBuf wbBuf

	// 如果为 1，表示在下一个安全点运行 sched.safePointFn 函数。
	//在 Golang 的运行时系统中，安全点是程序执行时的一种特殊状态，可以在此状态下进行垃圾回收和其他系统任务。
	//运行时系统会在安全点中暂停所有 Goroutine 的执行，并在安全点之间执行一些必要的操作，
	//例如垃圾回收的扫描和标记,调度器或其他系统任务相关的功能。
	runSafePointFn uint32 // if 1, run sched.safePointFn at next safe point

	//statsSeq是一个计数器，指示当前的 P 是否正在写入任何统计信息。
	//其值为奇数时表示正在写入统计信息时
	statsSeq atomic.Uint32

	//timersLock 是用于保护计时器（timers）的互斥锁（mutex）。
	//互斥锁是一种同步原语，它提供了对共享资源的独占访问权。通过在对计时器进行访问之前获取互斥锁，并在访问完成后释放锁，
	//可以确保同一时间只有一个 Goroutine 能够访问计时器，避免并发访问导致的数据竞争和不一致性。
	//在正常情况下，运行时系统会在运行在同一 P 上的 Goroutine 中访问计时器，但调度器也可能会在不同的 P 上访问计时器。
	timersLock mutex

	// timers 是一个用于存储定时器的数组，表示在某个特定时间要执行的操作。该字段用于实现标准库的 time 包。
	//访问 timers 字段时必须持有 timersLock 互斥锁,避免并发操作引起的竞态条件。
	timers []*timer

	// 表示 P（Processor） 的堆（heap）中的定时器数量。
	numTimers atomic.Uint32

	// 表示 P（Processor） 的堆（heap）中被删除的定时器数量
	deletedTimers atomic.Uint32

	// timerRaceCtx 字段用于在执行定时器函数时记录竞争上下文。
	//在并发环境下，多个 Goroutine 可能会同时访问和执行定时器函数，
	//因此需要使用竞争上下文来追踪和标识定时器函数的执行情况。
	timerRaceCtx uintptr

	//用于累积活跃 Goroutine（即可进行堆栈扫描的 Goroutine）所占用的堆栈空间大小。
	//在 Golang 的运行时系统中，垃圾回收器（GC）负责回收不再使用的内存。
	//堆栈扫描是垃圾回收的一部分，它用于识别堆栈上的对象并进行标记，以确保不会回收仍然被引用的对象。
	//maxStackScanDelta 字段用于跟踪并累积活跃 Goroutine 所占用的堆栈空间大小的变化。
	//当堆栈空间的变化达到一定阈值（maxStackScanSlack 或 -maxStackScanSlack），
	//maxStackScanDelta 字段的值会被刷新到 gcController.maxStackScan 字段中。
	//通过记录 maxStackScanDelta，运行时系统可以实时跟踪堆栈空间的使用情况， 并在达到阈值时触发相应的处理逻辑。
	//这有助于优化垃圾回收器的性能和效率，以及控制堆栈扫描的成本。
	maxStackScanDelta int64

	// scannedStackSize 和 scannedStacks 是用于记录与 GC 时间相关的有关当前 Goroutine 的统计信息的字段。
	//scannedStackSize 字段用于累积当前 P（Processor） 扫描的 Goroutine 堆栈的大小。
	//它表示在 GC 过程中扫描的 Goroutine 堆栈实际使用的空间大小（hi - sp）。
	//scannedStacks 字段用于累积当前 P 扫描的 Goroutine 数量。
	//这些字段的目的是跟踪当前 P 在 GC 过程中扫描的 Goroutine 的堆栈信息。
	//通过收集和记录这些统计数据，运行时系统可以评估 GC 过程中 Goroutine 堆栈的使用情况，
	//以提供对垃圾回收过程中 Goroutine 堆栈的细粒度监控和优化支持。
	scannedStackSize uint64 // stack size of goroutines scanned by this P
	scannedStacks    uint64 // number of goroutines scanned by this P

	// 用于指示该 P（Processor）是否应尽快进入调度器（scheduler），而不考虑当前运行在该 P 上的 Goroutine。
	//在 Golang 的运行时系统中，调度器负责协调 Goroutine 的执行。
	//为了实现公平的调度和避免 Goroutine 长时间占用 P，调度器会周期性地检查是否需要进行调度切换，即让其他 Goroutine 获取执行机会。
	//preempt 字段用于标记该 P 是否应立即进入调度器。
	//当 preempt 字段为 true 时，该 P 将优先进入调度器，即使当前正在运行的 Goroutine 尚未完成。
	//这可以有效地实现调度器的抢占式调度，以避免某个 Goroutine 长时间占用 P，导致其他 Goroutine 饥饿。
	//通过设置 preempt 字段，运行时系统可以实现对长时间运行的 Goroutine 的抢占，
	//确保其他 Goroutine 有机会获取执行时间，提高整体系统的公平性和性能。
	preempt bool

	// pageTraceBuf 是一个用于记录页面分配、释放和清理追踪的缓冲区。
	//在 Golang 的运行时系统中，可以启用 GOEXPERIMENT=pagetrace 实验特性来收集与页面管理相关的追踪信息。
	//当启用了 pagetrace 实验特性时，运行时系统会追踪页面的分配、释放和清理操作，并将相关的追踪信息记录下来。
	//pageTraceBuf 是一个用于缓存页面追踪信息的缓冲区。它会在需要记录页面追踪信息时被使用。
	//当收集到足够的追踪信息后，运行时系统会将其写入输出流或日志文件中，以供进一步分析和调试。
	//需要注意的是，pageTraceBuf 字段仅在启用了 pagetrace 实验特性时才会被使用。
	//它是运行时系统内部的一个工具，用于收集调试和性能分析所需的页面追踪信息。
	//通过使用 pageTraceBuf 字段，Golang 的运行时系统能够提供额外的工具和功能，以帮助开发人员诊断和优化与页面管理相关的问题。
	//它提供了对页面分配、释放和清理的细粒度追踪，有助于理解和分析运行时系统的内存管理行为。
	pageTraceBuf pageTraceBuf

	//pad cpu.CacheLinePad //注意：该字段在当前版本已被移除
	// p 结构体中的填充字段（pad）不再需要。填充字段在以前的版本中用于解决伪共享（False sharing）问题，
	//即在多个处理器上访问同一缓存行而导致的性能下降。
	//然而，在当前的实现中，p 结构体的大小已经足够大，以至于它的大小类是缓存行大小的整数倍（对于我们的任何架构都是如此）。
	//这意味着 p 结构体的大小已经足够大，不会与其他结构体或变量共享同一缓存行。
	//因此，填充字段不再需要来解决伪共享问题。这样一来，可以减少额外的填充字段，从而节省内存空间。
	//这个改变对于解决伪共享问题有一定的意义，因为它减少了不必要的内存开销，并提高了 p 结构体的紧凑性。
	//这使得每个 p 结构体在处理器之间的迁移和访问时更加高效，同时减少了缓存行的冲突，进一步提升了并发性能。
}
