package _defer

// 源码位置:$GOPATH/src/runtime/runtime2.go
type _defer struct {
	// 参数和返回值的内存大小
	siz int32

	//表示该_defer语句是否已经开始执行
	started bool

	//表示该_defer语句的优先级
	//当一个_defer语句被执行时，它会被添加到_defer链表中，而heap字段则用于将_defer语句添加到一个优先队列中，
	//以便在函数返回时按照一定的顺序执行_defer语句。在_defer链表中，后添加的_defer语句会先被执行，而在优先队列中，
	//heap值较小的_defer语句会先被执行。这个字段的值是在_defer语句被添加到_defer链表时根据一定规则计算出来的，
	//通常是根据_defer语句的执行顺序和作用域等因素计算而得。在函数返回时，Go语言会按照heap值的大小顺序执行_defer语句。
	//如果多个_defer语句的heap值相同，则它们会按照它们在_defer链表中的顺序依次执行。
	//这个机制可以确保_defer语句按照一定的顺序执行，从而避免了一些潜在的问题。
	heap bool

	// 表示该_defer用于具有开放式编码_defer的帧。开放式编码_defer是指在编译时已经确定_defer语句的数量和位置，
	//而不是在运行时动态添加_defer语句。在一个帧中，可能会有多个_defer语句，但只会有一个_defer结构体记录了所有_defer语句的信息，
	//而openDefer就是用来标识该_defer结构体是否是针对开放式编码_defer的
	openDefer bool

	//_defer语句所在栈帧的栈指针（stack pointer）
	//在函数调用时，每个函数都会创建一个新的栈帧，用于保存函数的局部变量、参数和返回值等信息。
	//而_defer语句也被保存在这个栈帧中，因此需要记录栈指针以便在函数返回时找到_defer语句。
	//当一个_defer语句被执行时，它会被添加到_defer链表中，并记录当前栈帧的栈指针。
	//在函数返回时，Go语言会遍历_defer链表，并执行其中的_defer语句。而在执行_defer语句时，
	//需要使用保存在_defer结构体中的栈指针来访问_defer语句所在栈帧中的局部变量和参数等信息。
	//需要注意的是，由于_defer语句是在函数返回之前执行的，因此在执行_defer语句时，函数的栈帧可能已经被销毁了。
	//因此，_sp字段的值不能直接使用，需要通过一些额外的处理来确保_defer语句能够正确地访问栈帧中的信息。
	sp uintptr

	//_defer语句的程序计数器（program counter）
	//程序计数器是一个指针，指向正在执行的函数中的下一条指令。在_defer语句被执行时，它会被添加到_defer链表中，
	//并记录当前函数的程序计数器。当函数返回时，Go语言会遍历_defer链表，并执行其中的_defer语句。
	//而在执行_defer语句时，需要让程序计数器指向_defer语句中的函数调用，以便正确地执行_defer语句中的代码。
	//这就是为什么_defer语句需要记录程序计数器的原因。需要注意的是，由于_defer语句是在函数返回之前执行的，
	//因此在执行_defer语句时，程序计数器可能已经指向了其它的函数或代码块。因此，在执行_defer语句时，
	//需要使用保存在_defer结构体中的程序计数器来确保_defer语句中的代码能够正确地执行。
	pc uintptr // pc 计数器值，程序计数器

	// defer 传入的函数地址，也就是延后执行的函数
	fn *funcval

	//defer 的 panic 结构体
	_panic *_panic

	//用于将多个defer链接起来，形成一个defer栈
	//当程序执行到一个 defer 语句时，会将该 defer 语句封装成一个 _defer 结构体，并将其插入到 defer 栈的顶部。
	//当函数返回时，程序会从 defer 栈的顶部开始依次执行每个 defer 语句，直到 defer 栈为空为止。
	//每个 _defer 结构体中的 link 字段指向下一个 _defer 结构体，从而将多个 _defer 结构体链接在一起。
	//当程序执行完一个 defer 语句后，会将该 defer 从 defer 栈中弹出，并将其 link 字段指向的下一个 _defer 结构体设置为当前的 defer 栈顶。
	//这样，当函数返回时，程序会依次执行每个 defer 语句，从而实现 defer 语句的反转执行顺序的效果。
	//需要注意的是，由于 _defer 结构体是在运行时动态创建的，因此 defer 栈的大小是不固定的。
	//在编写程序时，应该避免在单个函数中使用大量的 defer 语句，以免导致 defer 栈溢出。
	link *_defer
}



func deferproc(siz int32, fn *funcval) { // arguments of fn follow fn
	gp := getg() //获取goroutine结构
	if gp.m.curg != gp {
		// go code on the system stack can't defer
		throw("defer on system stack")
	}
	...
	d := newdefer(siz) //新建一个defer结构
	if d._panic != nil {
		throw("deferproc: d.panic != nil after newdefer")
	}
	d.link = gp._defer // 新建defer的link指针指向g的defer
	gp._defer = d      // 新建defer放到g的defer位置，完成插入链表表头操作
	d.fn = fn
	d.pc = callerpc
	d.sp = sp
	...
}
