package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/phper95/pkg/db"
	"log"
)

type User struct {
	ID   int64
	Name string
	Age  int64
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := db.InitMysqlClient(db.DefaultClient, "root", "admin123", "172.17.182.249:3306", "test")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	o := db.GetMysqlClient(db.DefaultClient)
	var sumTotal int64
	tableName := "user_big"
	if err := o.Raw("select count(*) from " + tableName).Count(&sumTotal); err != nil {
		log.Println(err)
	}
	log.Println("sumTotal", sumTotal)
	var minId int64
	limit := 10
	batchCnt := 0
	for {
		var users []User
		err := o.Raw("select * from "+tableName+" where id > ? order by id asc limit ? ", minId, limit).Scan(&users).Error
		if err != nil {
			log.Println(err)
			continue
		}
		if len(users) == 0 {
			log.Println("end!")
			return
		}
		batchCnt += 1
		minId = users[len(users)-1].ID
		for _, user := range users {
			if batchCnt <= 2 {
				log.Printf("%+v", user)
			}
		}

		if batchCnt <= 2 {
			log.Println("batchCnt", batchCnt)
		}
	}

}
