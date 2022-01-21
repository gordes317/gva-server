package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SelectNamespaces struct {
	Value string `json:"value" form:"value"`
	Label string `json:"label" form:"label"`
}

type SelectCluster struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Context string `json:"context" form:"context"`
}

type GetAllClusterInfo struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Context string `json:"context" form:"context"`
	Config  string `json:"config" form:"config"`
}

type MetricsCpuMem struct {
	Name string  `json:"name" form:"name"`
	Cpu  float32 `json:"cpu" form:"cpu"`
	Mem  float32 `json:"mem" form:"mem"`
}

type User struct {
	UserName string `json:"UserName" form:"UserName"`
}

type ClusterEvent struct {
	Reason        string      `json:"Reason" form:"Reason"`
	Namespace     string      `json:"Namespace" form:"Namespace"`
	Message       string      `json:"Message" form:"Message"`
	Resource      string      `json:"Resource" form:"Resource"`
	LastEventTime metav1.Time `json:"LastEventTime" form:"LastEventTime"`
}
