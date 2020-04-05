/**
* @Author: Chao
* @Date: 2020-04-04 19:49
* @Version: 1.0
 */

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo  --->  数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接Mysql数据库
	db, err := gorm.Open("mysql", "root:521chaochao@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := &UserInfo{1, "丁延超", "男", "蛙泳"}
	//db.Create(u1)

	//查询
	var u UserInfo
	db.First(&u)   //查询表中第一条数据保存到u中
	fmt.Printf("u:%#v\n", u)

	//更新
	db.Model(&u).Update("hobby", "双色球")

	//删除
	db.Delete(&u)
}
