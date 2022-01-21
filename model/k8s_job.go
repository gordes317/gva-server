package model

import (
	"time"
)

type Job struct {
	Name           string    `json:"name" form:"name"`
	Namespace      string    `json:"namespace" form:"namespace"`
	Successed      int32     `json:"successed" form:"successed"`
	Failed         int32     `json:"failed" form:"failed"`
	Completions    int32     `json:"completions" form:"completions"`
	StartTime      time.Time `json:"starttime" form:"starttime"`
	CompletionTime time.Time `json:"completiontime" form:"completiontime"`
	YamlData       string    `json:"yamldata" form:"yamldata"`
	CreateTime     time.Time `json:"createTime" form:"createTime"`
}

type JobUser struct {
	Job
	User
}
