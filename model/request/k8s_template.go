package request

import "gin-vue-admin/model"

type TemplateSearch struct {
	model.Template
	PageInfo
}
