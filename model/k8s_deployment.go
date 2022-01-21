package model

import (
	"time"
)

//specification for deployment from file
type Deployment struct {
	Name              string    `json:"name" form:"name"`
	Namespace         string    `json:"namespace" form:"namespace"`
	Replicas          int32     `json:"replicas" form:"replicas"`
	ReadyReplicas     int32     `json:"readyreplicas" form:"readyreplicas"`
	UpdateReplicas    int32     `json:"updatereplicas" form:"updatereplicas"`
	AvailableReplicas int32     `json:"availablereplicas" form:"availablereplicas"`
	Label             string    `json:"label" form:"label"`
	YamlData          string    `json:"yamldata" form:"yamldata"`
	CreateTime        time.Time `json:"createTime" form:"createTime"`
}

type DeploymentUser struct {
	Deployment
	User
}
