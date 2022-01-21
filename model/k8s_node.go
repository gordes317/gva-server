package model

import (
	"time"
)

type Node struct {
	Name           string    `json:"name" form:"name"`
	IP             string    `json:"ip" form:"ip"`
	Status         string    `json:"status" form:"status"`
	Role           string    `json:"role" form:"role"`
	Version        string    `json:"version" form:"version"`
	NodeSystem     string    `json:"nodesystem" form:"nodesystem"`
	KernelVersion  string    `json:"kernelversion" form:"kernelversion"`
	RunTime        string    `json:"runtime" form:"runtime"`
	PodsNumber     int32     `json:"podsnumber" form:"podsnumber"`
	CpuAllocatable float32   `json:"cpuallocatable" form:"cpuallocatable"`
	CpuTotal       float32   `json:"cputotal" form:"cputotal"`
	MemAllocatable float32   `json:"memallocatable" form:"memallocatable"`
	MemTotal       float32   `json:"memtotal" form:"memtotal"`
	Taint          string    `json:"taint" form:"taint"`
	CreateTime     time.Time `json:"createTime" form:"createTime"`
}

type NodeUser struct {
	Node
	User
}
