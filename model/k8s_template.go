// 自动生成模板Student
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Template struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(50);size:50;"`
	Type    string `json:"type" form:"type" gorm:"column:type;comment:类型;type:varchar(50);size:80;"`
	Data    string `json:"data" form:"data" gorm:"column:data;comment:内容;type:TEXT"`
	Comment string `json:"comment" form:"comment" gorm:"comment:nicname;comment:备注;type:varchar(100);size:100;"`
}

func (Template) TableName() string {
	return "template"
}
