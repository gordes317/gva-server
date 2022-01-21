package request

import "gin-vue-admin/model"

type StatefulSetSearch struct {
	model.StatefulSet
	PageInfo
}

type StatefulSetSearchUser struct {
	StatefulSetSearch
	model.User
}
