package model

import (
	"time"
)

type StorageClass struct {
	Name                 string `json:"name" form:"name"`
	Provisioner          string `json:"provisioner" form:"provisioner"`
	ReclaimPolicy        string `json:"reclaimpolicy" form:"reclaimpolicy"`
	VolumeBindingMode    string `json:"volumebindingmode" form:"volumebindingmode"`
	AllowVolumeExpansion bool   `json:"allowvolumeexpansion" form:"allowvolumeexpansion"`
	YamlData             string `json:"yamldata" form:"yamldata"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
}

type StorageClassUser struct {
	StorageClass
	User
}
