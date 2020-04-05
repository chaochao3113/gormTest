/**
* @Author: Chao
* @Date: 2020-04-05 16:13
* @Version: 1.0
 */

package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
//type User struct {
//	ID   int64 //gorm  默认将ID字段作为表的主键
//	Name *string  `gorm:"default:'小王子'"`
//	Age  int64
//}

type User struct {
	ID   int64 //gorm  默认将ID字段作为表的主键
	Name sql.NullString  `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	//连接Mysql数据库
	db, err := gorm.Open("mysql", "root:521chaochao@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3. 创建
	//user := &User{Age: 18}  //在代码层面创建一个User对象

	//user := &User{Name: "", Age: 48}  //设置了默认值的字段  再使用类型空值(比如 0, false, ""等等)将不起作用
	//解决方法:
	// (1) 将原结构体中的数据类型用指针表示
	//user := &User{Name: new(string), Age: 78}  //在代码层面创建一个User对象

	// (2) 使用实现了Scanner/Valuer接口的类
	user := &User{Name: sql.NullString{String: "", Valid: true}, Age: 88}  // Valid设置为true表示这是一个有值的字符串
	fmt.Println("主键是否为空", db.NewRecord(user))   //判断主键是否为空
	db.Debug().Create(user)
	fmt.Println("主键是否为空", db.NewRecord(user))   //判断主键是否为空

}
