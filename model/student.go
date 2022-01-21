// 自动生成模板Student
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Student struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;type:varchar(50);size:50;"`
      Age  int `json:"age" form:"age" gorm:"column:age;comment:年龄;type:int;size:10;"`
}


func (Student) TableName() string {
  return "student"
}

