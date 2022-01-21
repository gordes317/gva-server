package model

import (
	"time"
)

type Ingress struct {
	Name       string    `json:"name" form:"name"`
	Namespace  string    `json:"namespace" form:"namespace"`
	Host       string    `json:"host" form:"host"`
	Paths      string    `json:"paths" form:"paths"`
	YamlData   string    `json:"yamldata" form:"yamldata"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
}

type IngressUser struct {
	Ingress
	User
}
