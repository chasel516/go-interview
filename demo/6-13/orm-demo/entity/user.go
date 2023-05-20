package entity

import "time"

// 字段属性设置
type User struct {
	ID uint //bigint(20) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY
	//ID       uint64
	//UserID uint `gorm:"primarykey;AUTO_INCREMENT"` //自定义主键
	//Name   string
	Name string `gorm:"type:varchar(255);NOT NULL;DEFAULT:''"`
	//Name     *string
	Age      int `gorm:"type:tinyint(3);unsigned"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique"`
}
