package model

import (
	"time"
)

type DaemonSet struct {
	Name         string    `json:"name" form:"name"`
	Namespace    string    `json:"namespace" form:"namespace"`
	Desired      int32     `json:"desired" form:"desired"`
	Current      int32     `json:"current" form:"current"`
	Ready        int32     `json:"ready" form:"ready"`
	Updated      int32     `json:"updated" form:"updated"`
	Available    int32     `json:"available" form:"available"`
	NodeSelector string    `json:"nodeselector" form:"nodeselector"`
	YamlData     string    `json:"yamldata" form:"yamldata"`
	CreateTime   time.Time `json:"createTime" form:"createTime"`
}

type DaemonSetUser struct {
	DaemonSet
	User
}
