package model

import (
	"time"
)

type StatefulSet struct {
	Name       string    `json:"name" form:"name"`
	Namespace  string    `json:"namespace" form:"namespace"`
	Replicas   int32     `json:"replicas" form:"replicas"`
	Ready      int32     `json:"ready" form:"ready"`
	YamlData   string    `json:"yamldata" form:"yamldata"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
}

type StatefulSetUser struct {
	StatefulSet
	User
}
