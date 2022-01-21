package request

import "gin-vue-admin/model"

type DaemonSetSearch struct {
	model.DaemonSet
	PageInfo
}
type DaemonSetSearchUser struct {
	DaemonSetSearch
	model.User
}
