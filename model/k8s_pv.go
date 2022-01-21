package model

import (
	"time"
)

type PV struct {
	Name          string    `json:"name" form:"name"`
	Capacity      string    `json:"capacity" form:"capacity"`
	AccessMode    string    `json:"accessmode" form:"accessmode"`
	ReclaimPolicy string    `json:"reclaimpolicy" form:"reclaimpolicy"`
	Status        string    `json:"status" form:"status"`
	Claim         string    `json:"claim" form:"claim"`
	StorageClass  string    `json:"storageclass" form:"storageclass"`
	Reason        string    `json:"reason" form:"reason"`
	VolumeMode    string    `json:"volumemode" form:"volumemode"`
	YamlData      string    `json:"yamldata" form:"yamldata"`
	CreateTime    time.Time `json:"createTime" form:"createTime"`
}

type PVUser struct {
	PV
	User
}
