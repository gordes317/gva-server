package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)


type TektonSpec struct {
	TknSpec string `json:"TknSpec"`
	Image    string `json:"image"`
}

type Tekton struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TektonSpec `json:"spec,omitempty"`
}

type TektonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Tekton `json:"items"`
}

//type GroupVersionResource struct {
//	Group string
//	Version string
//	Resource string
//}

type TektonEvent struct {
	EventListener string `json:"EventListener" form:"EventListener"`
	EventListenerUID string `json:"EventListenerUID" form:"EventListenerUID"`
	EventID string `json:"EventID" form:"EventID"`
	TknNameSpace string `json:"TknNameSpace" form:"TknNameSpace"`
}


type TektonEventUser struct {
	TektonEvent
	User
}


type TektonPipelineRunsList struct {
	Name          string      `json:"name" form:"name"`
	Namespace     string      `json:"namespace" form:"namespace"`
	Pipeline     string      `json:"pipeline" form:"pipeline"`
	Message      string      `json:"message" form:"message"`
	Reason      string      `json:"reason" form:"reason"`
	Status      string      `json:"status" form:"status"`
	Watch      bool      `json:"watch" form:"watch"`
	TriggersEventId      string      `json:"triggersEventId" form:"triggersEventId"`
	CompletionTime    *metav1.Time   `json:"completionTime" form:"completionTime"`
	LastTransitionTime    metav1.Time   `json:"lastTransitionTime" form:"lastTransitionTime"`
	StartTime    *metav1.Time   `json:"startTime" form:"startTime"`
	CreateTime    time.Time   `json:"createTime" form:"createTime"`
}


type TektonPipeline struct {
	Name          string      `json:"name" form:"name"`
	Namespace     string      `json:"namespace" form:"namespace"`
	Tasks     []Tasks      `json:"tasks" form:"tasks"`
}


type Tasks struct {
	TaskName string `json:"taskName" form:"taskname"`
	TaskRefName string `json:"taskRefName" form:"taskRefName"`
}


type TektonPipelineTaskList struct {
	Name          string      `json:"taskName" form:"taskName"`
	Namespace     string      `json:"namespace" form:"namespace"`
	PodName     string      `json:"podName" form:"podName"`
	Reason     string      `json:"reason" form:"reason"`
	Steps     []TaskStep      `json:"steps" form:"steps"`
}

type Step struct {
	StepName string`json:"stepName" form:"stepName"`
}


type TektonTaskRunList struct {
	Name          string      `json:"name" form:"name"`
	Namespace     string      `json:"namespace" form:"namespace"`
	PipelineRunName     string      `json:"pipelineRunName" form:"pipelineRunName"`
	PodName     string      `json:"podName" form:"podName"`
	Reason     string      `json:"reason" form:"reason"`
	Status     string      `json:"status" form:"status"`
	Message     string      `json:"message" form:"message"`
	TaskRefName     string      `json:"taskName" form:"taskName"`
	Steps     []TaskStep      `json:"steps" form:"steps"`
}

type TaskStep struct {
	StepName string `json:"stepName" form:"stepName"`
	StepContainerName string `json:"stepContainerName" form:"stepContainerName"`
	StartAt metav1.Time `json:"startAt" form:"startAt"`
	FinishAt metav1.Time `json:"finishAt" form:"finishAt"`
	Reason string `json:"reason" form:"reason"`
}