package model

import (
	"time"
)

// Service specification for deployment from file
type Service struct {
	Name        string    `json:"name" form:"name"`
	Namespace   string    `json:"namespace" form:"namespace"`
	Type        string    `json:"type" form:"type"`
	Ports       string    `json:"ports" form:"ports"`
	ClusterIP   string    `json:"clusterip" form:"clusterip"`
	ExternalIPs string    `json:"externalips" form:"externalips"`
	Selector    string    `json:"selector" form:"selector"`
	YamlData    string    `json:"yamldata" form:"yamldata"`
	CreateTime  time.Time `json:"createTime" form:"createTime"`
}

type ServiceUser struct {
	Service
	User
}
