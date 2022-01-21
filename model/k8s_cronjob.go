package model

import (
	"time"
)

type CronJob struct {
	Name             string    `json:"name" form:"name"`
	Namespace        string    `json:"namespace" form:"namespace"`
	Schedule         string    `json:"schedule" form:"schedule"`
	Suspend          bool      `json:"suspend" form:"suspend"`
	Active           string    `json:"avtive" form:"avtive"`
	LastScheduleTime time.Time `json:"lastscheduletime" form:"lastscheduletime"`
	YamlData         string    `json:"yamldata" form:"yamldata"`
	CreateTime       time.Time `json:"createTime" form:"createTime"`
}

type CronJobUser struct {
	CronJob
	User
}
