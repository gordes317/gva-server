package model

import (
	"time"
)

type ConfigMap struct {
	Name       string            `json:"name" form:"name"`
	Namespace  string            `json:"namespace" form:"namespace"`
	Data       map[string]string `json:"data" form:"data"`
	YamlData   string            `json:"yamldata" form:"yamldata"`
	CreateTime time.Time         `json:"createTime" form:"createTime"`
}

type ConfigMapUser struct {
	ConfigMap
	User
}
