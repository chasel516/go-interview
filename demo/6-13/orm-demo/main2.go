package main

import (
	"errors"
	"github.com/phper95/pkg/db"
	"gorm.io/gorm"
	"log"
	"orm-demo/entity"
)

const DBName = "test"

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initDB()
}

func initDB() {
	err := db.InitMysqlClient(db.DefaultClient, "root", "admin123", "localhost:3306", DBName)
	if err != nil {
		log.Fatal("InitMysqlClient client error" + db.DefaultClient)
	}
	log.Print("connect mysql success ", db.DefaultClient)
	err = db.InitMysqlClientWithOptions(db.TxClient, "root", "admin123", "localhost:3306", DBName, db.WithPrepareStmt(false))
	if err != nil {
		log.Print("InitMysqlClient client error" + db.TxClient)
		return
	}
}

func main() {
	//此实例不支持事务
	//ormDB := db.GetMysqlClient(db.DefaultClient).DB

	//需要关闭SQL预编译才能使用事务
	ormDBTx := db.GetMysqlClient(db.TxClient).DB

	//建表
	if err := ormDBTx.AutoMigrate(&entity.User{}); err != nil {
		log.Print("AutoMigrate user error", err)
	}

	//定义3行数据
	user1 := entity.User{
		Name:     "user1",
		Age:      0,
		Birthday: nil,
		Email:    "user1@qq.com",
	}

	user2 := entity.User{
		Name:     "user2",
		Age:      0,
		Birthday: nil,
		Email:    "user2@qq.com",
	}

	user3 := entity.User{
		Name:     "user3",
		Age:      0,
		Birthday: nil,
		Email:    "user3@qq.com",
	}
	//嵌套事务
	//通过Transaction方法传入一个回调函数，在回调函数中执行一系列数据库操作，如果回调函数返回错误，则回滚事务，否则提交事务
	err := ormDBTx.Transaction(func(tx *gorm.DB) error {
		//注意，内部需要使用tx，而不是ormDB
		tx.Create(&user1)

		err := tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&user2)
			return errors.New("rollback user2") // 回滚 user2
		})
		if err != nil {
			log.Print("Create user2 error", err)
		}

		err = tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&user3)
			return nil
		})
		if err != nil {
			log.Print("Create user3 error", err)
		}

		//返回值为nil时才会提交事务
		return nil
	})
	if err != nil {
		log.Print("Transaction error", err)
	}

	user4 := entity.User{
		Name:     "user4",
		Age:      0,
		Birthday: nil,
		Email:    "user4@qq.com",
	}

	user5 := entity.User{
		Name:     "user5",
		Age:      0,
		Birthday: nil,
		Email:    "user5@qq.com",
	}

	//通过Begin开启事务之后会返回一个事务操作对象*gdb.TX
	tx := ormDBTx.Begin()
	tx.Create(&user4)

	//保存点
	tx.SavePoint("step1")
	tx.Create(&user5)

	//回滚所有操作
	//tx.Rollback()
	//回滚到保存点
	tx.RollbackTo("step1") // 回滚 user2
	tx.Commit()            // 最终仅提交 user4

}
