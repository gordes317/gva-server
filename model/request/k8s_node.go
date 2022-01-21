package request

import "gin-vue-admin/model"

type NodeSearch struct {
	model.Node
	PageInfo
}

type NodeSearchUser struct {
	NodeSearch
	model.User
}

