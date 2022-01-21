package model

import (
	"time"
)

type Namespace struct {
	Name       string    `json:"name" form:"name"`
	Status     string    `json:"status" form:"status"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
}

type NamespaceUser struct {
	Namespace
	OldName string
	User
}
