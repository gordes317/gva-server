package request

import "gin-vue-admin/model"

type TektonPipelineRunsListSearch struct {
	model.TektonPipelineRunsList
	PageInfo
}
type TektonPipelineRunsListSearchUser struct {
	TektonPipelineRunsListSearch
	model.User
}

type TektonPipelineSearch struct {
	model.TektonPipeline
	PageInfo
}
type TektonPipelineSearchUser struct {
	TektonPipelineSearch
	model.User
}

type TektonPipelineTaskUser struct {
	model.TektonPipelineTaskList
	model.User
}

type TektonTaskRunListUser struct {
	model.TektonTaskRunList
	model.User
}


type SearchBuildPlanParams struct {
	model.SysBuildPlan
	PageInfo
}
