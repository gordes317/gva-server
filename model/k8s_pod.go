package model

import (
	"time"
)

type Container struct {
	CName  string `json:"cname" form:"cname"`
	CImage string `json:"cimage" form:"cimage"`
	CPort  string `json:"cport" form:"cport"`
	//CState        string `json:"cstate" form:"cstate"`
	CRestartCount int32 `json:"crestartcount" form:"crestartcount"`
	CReason         string `json:"creason" form:"creason"`
}

type Pod struct {
	Name          string      `json:"name" form:"name"`
	Namespace     string      `json:"namespace" form:"namespace"`
	ContainerName string      `json:"containername" form:"containername"`
	Phase         string      `json:"phase" form:"phase"`
	Status         string      `json:"status" form:"status"`
	HostIP        string      `json:"hostip" form:"hostip"`
	PodIP         string      `json:"podip" form:"podip"`
	CPU           float32     `json:"cpu" form:"cpu"`
	Mem           float32     `json:"mem" form:"mem"`
	RestartCount  int32       `json:"restartcount" form:"restartcount"`
	Image         string      `json:"image" form:"image"`
	Log           string      `json:"log" form:"log"`
	Containers    []Container `json:"ContainerList"`
	YamlData      string      `json:"yamldata" form:"yamldata"`
	CreateTime    time.Time   `json:"createTime" form:"createTime"`
}


type PodUser struct {
	Pod
	User
}
