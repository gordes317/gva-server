// 自动生成模板Student
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type MultiCluster struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:集群名字;type:varchar(50);size:50;"`
	Address string `json:"address" form:"address" gorm:"column:address;comment:MasterUrl;type:varchar(200);size:200;"`
	Context string `json:"context" form:"context" gorm:"column:context;comment:上下文名称;type:varchar(200);size:200;"`
	Config  string `json:"config" form:"config" gorm:"column:config;comment:配置;type:TEXT"`
	Comment string `json:"comment" form:"comment" gorm:"comment:nicname;comment:备注;type:varchar(100);size:100;"`
}

func (MultiCluster) TableName() string {
	return "multicluster"
}
