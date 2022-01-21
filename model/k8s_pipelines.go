package model

import (
	"gin-vue-admin/global"
)

type SysBuildPlan struct {
	global.GVA_MODEL
	PlanName        string `json:"planName" gorm:"comment:构建计划名称"`
	RepoId string `json:"repoId" gorm:"comment:构建代码仓库ID"`
	RepoName string `json:"repoName" gorm:"comment:构建代码仓库名称"`
	ImageUrl    string `json:"imageUrl" gorm:"comment:镜像仓库地址"`
	ProjectName    string `json:"projectName" gorm:"comment:项目名称"`
	ImageName      string `json:"imageName" gorm:"comment:镜像名称"`
}
