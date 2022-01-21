package request

import "gin-vue-admin/model"

type PodSearch struct {
	model.Pod
	PageInfo
}
type PodSearchUser struct {
	PodSearch
	model.User
}
