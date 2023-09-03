package main

import (
	"github.com/phper95/pkg/db"
	"log"
)

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initMysql()
}

func initMysql() {
	err := db.InitMysqlClient(db.DefaultClient, "root", "admin123", "localhost:3306", "shop")
	if err != nil {
		log.Fatal("InitMysqlClient client error" + db.DefaultClient)
	}
	log.Print("connect mysql success ", db.DefaultClient)
	err = db.InitMysqlClientWithOptions(db.TxClient, "root", "admin123", "localhost:3306", "shop", db.WithPrepareStmt(false))
	if err != nil {
		log.Print("InitMysqlClient client error" + db.TxClient)
		return
	}

}

//type entity.User struct-demo {
//	gorm.Model
//	Name     string
//	Age      int `gorm:"type:tinyint(3);unsigned"`
//	Birthday *time.Time
//	Email    string `gorm:"type:varchar(100);unique"`
//}

//	type Tabler interface {
//		TableName() string
//	}
//
// // TableName 会将 User 的表名重写为 `user_table1`
//
//	func (entity.User) TableName() string {
//		return "user_table1"
//	}
func main() {

	//ormDB := db.GetMysqlClient(db.DefaultClient).DB

	//查看连接配置
	//sqlDB, _ := ormDB.DB()
	//log.Printf("Stats : %+v", sqlDB.Stats())

	//建表
	//if err := ormDB.AutoMigrate(&entity.User{}); err != nil {
	//	log.Print("AutoMigrate user error", err)
	//}

	//自定义表名的另一种方式
	//if err := ormDB.Table("user_table2").AutoMigrate(&entity.User{}); err != nil {
	//	log.Print("AutoMigrate user error", zap.Error(err))
	//}

	//name := ""
	//写入数据
	//user := entity.User{
	//	Name: name,
	//	//Name:     &name,
	//	Age:      0,
	//	Birthday: nil,
	//	Email:    "111@qq.com",
	//}
	//if err := ormDB.Create(&user).Error; err != nil {
	//	log.Print("insert error", zap.Any("user", user))
	//}

	// 指定字段创建
	//if err := ormDB.Select("email").Create(&user).Error; err != nil {
	//	log.Print("insert error", zap.Any("user", user))
	//}

	// 批量创建
	//var users = []entity.User{{Name: "user1", Email: "u1"}, {Name: "user2", Email: "u2"}, {Name: "user3", Email: "u3"}}
	//if err := ormDB.Create(&users).Error; err != nil {
	//	log.Print("insert error", zap.Any("user", user))
	//}

	//查询时会忽略空值，o值，false和null值
	//users := make([]entity.User, 0)
	//ormDB.Where(&user).Find(&users)
	//log.Printf("%+v", users)

	//只获取指定字段
	//方式1
	//ormDB.Select("name", "email").Find(&users)
	//log.Print("指定字段", users)
	//方式2 通过定义结构体来限制需要获取的字段
	//type APIUser struct-demo {
	//	ID   uint
	//	Name string
	//}
	//ormDB.Model(&entity. {}).Limit(5).Find(&APIUser{}).Scan(&users)
	//
	//执行原生sql
	//查询
	//var userRes entity.User
	//err := ormDB.Raw("select * from user where id = ?", 1).Scan(&userRes).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(userRes)
	//
	//err = ormDB.Exec("DROP TABLE user").Error
	//if err != nil {
	//	log.Print("DROP TABLE error", err)
	//}
}
