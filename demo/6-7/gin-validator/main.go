package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type User struct {
	//gin通过binding标签来定义校验规则
	//required 表示该参数必须存在
	Name string `json:"name" binding:"required"`

	//gte=18表示该参数的值必须大于等于18
	Age int `json:"age" binding:"gte=18"`

	//age值等于18
	//Age      int    `json:"age" binding:"eq=18"`

	//age值不等于18
	//Age      int    `json:"age" binding:"ne=18"`

	//email 合法的邮箱格式
	Email string `json:"email" binding:"required,email"`

	//Password字段长度至少为6位且包含123@
	Password string `json:"password" binding:"required,contains=123@,min=6"`

	//eqfield表示跟指定字段相等，同理还有 qfield ,nefield ,gtfield ,gtefield ,ltfield ,ltefield
	//注意，这里的字段需要跟结构体中的字段名一致，而非是结构体字段的标签名
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`

	//url 合法的url格式
	//URL string `json:"url" binding:"required,url"`

	//合法的ip格式
	//IP string `json:"ip" binding:"required,ip"`

	//合法的ipv4格式
	//IPv4 string `json:"ipv4" binding:"required,ipv4"`

	//合法的ipv6格式
	//IPv6 string `json:"ipv6" binding:"required,ipv6"`

	//[]string长度必须大于1，数组中元素string长度必须在2-100之间
	Tags []string `json:"tags" binding:"gt=1,dive,required,min=2,max=100"`

	//keys 和 endkeys来标记key值的校验范围，从keys开始，至endkeys结束
	//限制key值长度必须在2-100之间，value值的长度也必须在2-100之间
	M      map[string]string `json:"m" binding:"dive,keys,min=2,max=100,endkeys,required,min=2,max=100"`
	Extra1 st                `json:"extra1"`
	Extra2 st                `json:"extra2" binding:"structonly"`
	Extra3 st                `json:"extra3" binding:"-"`

	//自定义验证器：dateLteNow,时间戳必须小于等于当前时间
	Mtime int64 `json:"mtime" binding:"dateLteNow"`
}

type st struct {
	F1 string `json:"f1" binding:"required,min=6"`
}

func main() {
	router := gin.Default()
	//注册验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//绑定第一个参数是验证的函数第二个参数是自定义的验证函数
		v.RegisterValidation("dateLteNow", dateLteNow)
	}
	router.POST("register", Register)
	router.Run(":8888")
}
func Register(c *gin.Context) {
	var u User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		fmt.Println("param check failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	//验证 存储操作省略.....
	fmt.Println("register success")
	c.JSON(http.StatusOK, "successful")
}

// 自定义验证函数，通过反射获取字段值并进行校验
func dateLteNow(fileLevel validator.FieldLevel) bool {
	//获取字段值
	t := fileLevel.Field().Int()
	fmt.Println("t:", t)
	if t == 0 {
		return false
	}
	//与当前时间对比
	if time.Now().Unix()-t < 0 {
		return false
	}
	return true
}
