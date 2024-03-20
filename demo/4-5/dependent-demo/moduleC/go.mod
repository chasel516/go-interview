module moduleC

go 1.20

require (
	gitee.com/phper95/pkg/errors v0.0.0-20231128051225-009aeeada524
	github.com/astaxie/beego v1.12.3
)

//moduleC 依赖gitee.com/phper95/pkg/errors，而gitee.com/phper95/pkg/errors依赖github.com/pkg/errors和go.uber.org/zap
//由于github.com/pkg/errors没有go.mod，所以github.com/pkg/errors被作为间接依赖导入
require github.com/pkg/errors v0.9.1 // indirect

//moduleC依赖beego,而beego依赖ansicolor,但由于ansicolor没有go.mod，所以ansicolor作为间接依赖添加到了moduleC的go.mod文件中
require github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
