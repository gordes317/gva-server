package utils

import (
	"gin-vue-admin/model"
)

type MapSlice []interface{}

func (m MapSlice) Len() int {
	return len(m)
}

func (m MapSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MapSlice) Less(i, j int) bool {
	m1 := m[i].(model.ClusterEvent).LastEventTime.Unix()
	m2 := m[j].(model.ClusterEvent).LastEventTime.Unix()

	return m1 > m2
}

type PipelineRunSlice []interface{}

func (m PipelineRunSlice) Len() int {
	return len(m)
}

func (m PipelineRunSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m PipelineRunSlice) Less(i, j int) bool {
	m1 := m[i].(model.TektonPipelineRunsList).CreateTime.Unix()
	m2 := m[j].(model.TektonPipelineRunsList).CreateTime.Unix()

	return m1 > m2
}
