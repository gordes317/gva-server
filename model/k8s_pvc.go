package model

import (
	"time"
)

type PVC struct {
	Name         string    `json:"name" form:"name"`
	Namespace    string    `json:"namespace" form:"namespace"`
	Status       string    `json:"status" form:"status"`
	Volume       string    `json:"volume" form:"volume"`
	Capacity     string    `json:"capacity" form:"capacity"`
	AccessMode   string    `json:"accessmode" form:"accessmode"`
	StorageClass string    `json:"storageclass" form:"storageclass"`
	VolumeMode   string    `json:"volumemode" form:"volumemode"`
	YamlData     string    `json:"yamldata" form:"yamldata"`
	CreateTime   time.Time `json:"createTime" form:"createTime"`
}

type PVCUser struct {
	PVC
	User
}

