module moduleA

go 1.20

require (
	//moduleA依赖moduleB，而moduleB没有go.mod文件，但moduleB依赖beego，beego依赖ansicolor
	//moduleB没有go.mod文件时，moduleA会将moduleB的直接依赖和间接依赖作为moduleA的间接依赖导入
	github.com/astaxie/beego v1.12.3 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
)
