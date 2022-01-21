package request

import "gin-vue-admin/model"

type MultiClusterSearch struct {
	model.MultiCluster
	PageInfo
}
