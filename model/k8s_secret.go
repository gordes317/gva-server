package model

import (
	"time"
)

type Secret struct {
	Name       string            `json:"name" form:"name"`
	Namespace  string            `json:"namespace" form:"namespace"`
	Type       string            `json:"type" form:"type"`
	Data       map[string][]byte `json:"data" form:"data"`
	YamlData   string            `json:"yamldata" form:"yamldata"`
	CreateTime time.Time         `json:"createTime" form:"createTime"`
}

type SecretUser struct {
	Secret
	User
}
