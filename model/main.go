/**
* @Author: Chao
* @Date: 2020-04-05 15:30
* @Version: 1.0
 */

package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model    //内嵌的gorm.Model
	Name         string
	Age          sql.NullInt64  `gorm:"column:user_age"`  // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {     // 创建到mysql中默认使用复数形式  ----> animals   也可以自己设置
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

//唯一指定表名
func (Animal) TableName() string {
	return "dyc"
}

func main() {
	//设置默认的表名前缀  修改默认的表名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_" + defaultTableName
	}

	//连接Mysql数据库
	db, err := gorm.Open("mysql", "root:521chaochao@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)   //禁用表名默认负数

	//自动迁移  相当于创建表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	//使用User结构体 创建名叫xiaowangzi的表
	//db.Table("xiaowangzi").CreateTable(&User{})
}
