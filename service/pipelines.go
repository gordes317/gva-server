package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/xanzy/go-gitlab"
	"gorm.io/gorm"
	"io/ioutil"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

//@author: jun.huang
//@function: StartCi
//@description: 开始Ci流水线
//@param: repository model.Repository
//@return: err error map[string]string

func StartCi(repository model.Repository) (err error, dat map[string]string) {

	client := &http.Client{}
	//var branchName string
	var branchOrTag string
	var preCommit string
	var commitId string
	if repository.CommitShortId == "1" {
		branchOrTag = "Branch"
		preCommit = repository.BranchName + "-"
	} else if repository.CommitShortId == "2" {
		branchOrTag = "Tag"
		preCommit = repository.BranchName + "-"
	} else {
		branchOrTag = "Tag"
		//preCommit = repository.BranchName + "-"
	}

	err2, commitIdMap := FindRepoCommitId(model.Branch{
		PId:         repository.PId,
		BranchName:  repository.BranchName,
		TagName:  repository.BranchName,
		BranchOrTag: branchOrTag,
	})

	if err2 != nil {
		panic(err)
	}
	if branchOrTag == "Tag" {
		commitId = repository.BranchName
	} else {
		commitId = preCommit + commitIdMap["commitId"]
	}

	data := make(map[string]interface{})
	//data["checkout_sha"] = repository.CommitShortId
	data["checkout_sha"] = commitId
	data["env"] = repository.Env
	repo := make(map[string]interface{})
	repo["git_ssh_url"] = "git@27.115.15.12:"+repository.ProjectName+".git"

	repo["branch_name"] = repository.BranchName
	//repo["name"] = repository.ProjectName
	ProjectName := strings.Split(repository.ProjectName, "/")[1]
	repo["name"] = ProjectName
	data["repository"] = repo
	fmt.Println("data:", data)
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST","http://tekton-listener.com/", bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-GitLab-Token", "1234567")
	req.Header.Set("X-Gitlab-Event", "Push Hook")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//var dat map[string]string
	err = json.Unmarshal(body, &dat)
	if err != nil {
		panic(err)
	}

	return err, dat
}

//@author: jun.huang
//@function: GetRepositoryList
//@description: 获取gitlab项目列表
//@param: Job model.Repository
//@return: err error []interface{}

func GetRepositoryProjectList() (err error, projectList []map[string]string) {

	gitlabUrl := "http://172.16.25.252/"
	accessToken := "u_siAHfRPegfHnWPezQ6"
	gl, _ := gitlab.NewClient(accessToken, gitlab.WithBaseURL(gitlabUrl))
	applications, _, err := gl.Projects.ListProjects(&gitlab.ListProjectsOptions{
		ListOptions:              gitlab.ListOptions{PerPage: 50000},
	})
	if err != nil {
		log.Fatal(err)
	}
	//var projectMap map[string]interface{}
	projectList = make([]map[string]string, 0)
	for _, p := range applications {
		var projectOption map[string]string
		projectOption = make(map[string]string)
		projectOption["value"] = strconv.Itoa(p.ID)
		projectOption["label"] = p.Namespace.Name + "/" +p.Name
		projectList = append(projectList, projectOption)
	}

	return err, projectList
}

//@author: jun.huang
//@function: FindRepoBranches
//@description: 获取gitlab项目分支和Tag
//@param: Job model.Repository
//@return: err error []interface{}

func FindRepoBranches(PId string) (err error, branchOptions []map[string]interface{}) {

	gitlabUrl := "http://172.16.25.252/"
	accessToken := "u_siAHfRPegfHnWPezQ6"
	gl, _ := gitlab.NewClient(accessToken, gitlab.WithBaseURL(gitlabUrl))

	branchList, _, err := gl.Branches.ListBranches(PId, &gitlab.ListBranchesOptions{})

	if err != nil {
		log.Fatal(err)
	}
	tagList, _, err := gl.Tags.ListTags(PId, &gitlab.ListTagsOptions{})

	if err != nil {
		log.Fatal(err)
	}
	branchesLabelArr := [2]string{"Branches", "Tags"}
	branchOptions = make([]map[string]interface{}, 0)
	for _, v := range branchesLabelArr {
		var subBranchLabelOptions map[string]interface{}
		subBranchLabelOptions = make(map[string]interface{})
		var subBranchesValueOptions []map[string]interface{}
		subBranchesValueOptions = make([]map[string]interface{}, 0)
		if v == "Branches" {
			subBranchLabelOptions["label"] = "Branches"
			for _, branch := range branchList {
				var realBranchesValueOptions map[string]interface{}
				realBranchesValueOptions = make(map[string]interface{})
				realBranchesValueOptions["value"] = branch.Name
				realBranchesValueOptions["Label"] = branch.Name
				subBranchesValueOptions = append(subBranchesValueOptions, realBranchesValueOptions)
			}
			subBranchLabelOptions["options"] = subBranchesValueOptions
			branchOptions = append(branchOptions, subBranchLabelOptions)
		} else{
			subBranchLabelOptions["label"] = "Tags"
			for _, tag := range tagList {
				//subBranchLabelOptions["label"] = "Branches"
				var realBranchesValueOptions map[string]interface{}
				realBranchesValueOptions = make(map[string]interface{})
				realBranchesValueOptions["value"] = tag.Name
				realBranchesValueOptions["Label"] = tag.Name
				subBranchesValueOptions = append(subBranchesValueOptions, realBranchesValueOptions)
			}
			subBranchLabelOptions["options"] = subBranchesValueOptions
			branchOptions = append(branchOptions, subBranchLabelOptions)
		}
	}

	return err, branchOptions
}


//@author: jun.huang 教程：https://mozillazg.com/2020/07/k8s-kubernetes-client-go-list-get-create-update-patch-delete-crd-resource-without-generate-client-code-update-or-create-via-yaml.html
//@function: FindRepoCommitId
//@description: 根据项目ID和分支名称或者tag名称获取项目CommitId
//@param: Pid model.Branch
//@return: err error map[string]string

func FindRepoCommitId(b model.Branch) (err error, commitIdMap map[string]string) {
	//fmt.Println("b:", b)
	gitlabUrl := "http://172.16.25.252/"
	accessToken := "u_siAHfRPegfHnWPezQ6"
	gl, _ := gitlab.NewClient(accessToken, gitlab.WithBaseURL(gitlabUrl))
	var commitShortId string
	var commitLongId string
	if b.BranchOrTag == "Branch" {
		branch, _, err := gl.Branches.GetBranch(b.PId, b.BranchName)
		if err != nil {
			panic(err)
		}
		commitShortId = branch.Commit.ShortID
		commitLongId = branch.Commit.ID
	} else {
		tag, _, err := gl.Tags.GetTag(b.PId, b.TagName)
		if err != nil {
			panic(err)
		}
		commitShortId = tag.Commit.ShortID
		commitLongId = tag.Commit.ID
	}
	commitIdMap = make(map[string]string)
	commitIdMap["commitShortId"] = commitShortId
	commitIdMap["commitId"] = commitLongId

	return err, commitIdMap
}


// GetTektonPipelineRunsList
// @Tags pipelines
// @Summary 分页获取Tekton PipelineRuns 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TektonPipelineRunsSearchUser true "分页获取Tekton PipelineRuns 列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelines/getTektonPipelineRunsList [get]

func GetTektonPipelineRunsList(tkn request.TektonPipelineRunsListSearchUser) (err error, list interface{}, total int64) {
	config := global.GetK8sConfig(tkn.UserName)
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return
	}
	//fmt.Printf("watch:%v\n", tkn.Watch)
	result, err := clientset.TektonV1alpha1().PipelineRuns(tkn.Namespace).List(context.TODO(), v1.ListOptions{Limit: 5000, Watch: tkn.Watch})
	//clientset.TektonV1alpha1().PipelineRuns(tkn.Namespace).Watch()
	//fmt.Printf("watch:%v\n,result:%v\n", tkn.Watch, result)
	if err != nil {
		return
	}
	var pipelineRuns model.TektonPipelineRunsList
	listResult := make([]interface{}, 0)
	//Items := result.Items
	//sort.Sort(utils.MapSlice(Items))
	for _, pipelineRun := range result.Items {
		pipelineRuns.Name = pipelineRun.Name
		pipelineRuns.Namespace = pipelineRun.Namespace
		pipelineRuns.Pipeline = pipelineRun.Spec.PipelineRef.Name
		pipelineRuns.TriggersEventId = pipelineRun.Labels["triggers.tekton.dev/triggers-eventid"]
		pipelineRuns.Message = pipelineRun.Status.Conditions[0].Message
		pipelineRuns.Reason = pipelineRun.Status.Conditions[0].Reason
		pipelineRuns.Status = string(pipelineRun.Status.Conditions[0].Status)
		pipelineRuns.CompletionTime = pipelineRun.Status.CompletionTime
		if pipelineRun.Status.CompletionTime == nil {
			pipelineRuns.CompletionTime = &v1.Time{Time: time.Now()}
		}
		//pipelineRuns.CompletionTime = pipelineRun.Status.CompletionTime
		pipelineRuns.LastTransitionTime = pipelineRun.Status.Conditions[0].LastTransitionTime.Inner
		pipelineRuns.StartTime = pipelineRun.Status.StartTime
		pipelineRuns.CreateTime = pipelineRun.CreationTimestamp.Time

		if strings.Contains(pipelineRun.Name, tkn.Name) && strings.Contains(pipelineRun.Namespace, tkn.Namespace) {
			listResult = append(listResult, pipelineRuns)
		}
	}

	sort.Sort(utils.PipelineRunSlice(listResult))

	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(tkn.Page, tkn.PageSize, listResult)

	return err, list, total

}


// GetPipeline
// @Tags Pipelines
// @Summary 用pipelineName查询Tasks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "用pipelineName查询Tasks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/getPipelineTasks [get]
func GetPipeline(name string, namespace string, userName string) (err error, pipeline model.TektonPipeline){

	// create the clientset
	clientset, err := versioned.NewForConfig(global.GetK8sConfig(userName))
	if err != nil {
		panic(err.Error())
	}

	result, err := clientset.TektonV1alpha1().Pipelines(namespace).Get(context.TODO(), name, v1.GetOptions{})
	if err != nil {
		return
	}
	
	//var pipeline model.TektonPipeline
	pipeline.Name = result.Name
	pipeline.Namespace = result.Namespace
	var tasks []model.Tasks
	var _task model.Tasks
	for _, t := range result.Spec.Tasks {
		_task.TaskName = t.Name
		_task.TaskRefName = t.TaskRef.Name
		tasks = append(tasks, _task)
	}
	pipeline.Tasks = tasks

	return err, pipeline
}


// GetPipelineTaskList
// @Tags pipelines
// @Summary 获取pipeline task列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TektonPipelineUser true "获取pipeline task列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelines/getPipelineTaskList [get]

func GetPipelineTaskList(name, namespace, userName string, page, pageSize int) (err error, list interface{}, total int64) {
	config := global.GetK8sConfig(userName)
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return
	}
	result, err := clientset.TektonV1alpha1().Tasks(namespace).List(context.TODO(), v1.ListOptions{Limit: 5000})
	if err != nil {
		return
	}
	var tasks model.TektonPipelineTaskList

	//listResult = []model.TektonPipelineTaskList
	listResult := make([]interface{}, 0)

	for _, t := range result.Items {
		tasks.Name = t.Name
		tasks.Namespace = t.Namespace
		tasks.Reason = "Waiting"
		var step model.TaskStep
		var stepList []model.TaskStep
		for _, s := range t.Spec.Steps {
			step.StepName = s.Name
			step.Reason = "Waiting"
			stepList = append(stepList, step)
		}

		tasks.Steps = stepList

		if strings.Contains(t.Name, name) && strings.Contains(t.Namespace, namespace) {
			listResult = append(listResult, tasks)
		}
	}

	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(page, pageSize, listResult)

	return err, list, total

}


// GetTaskRunList
// @Tags Pipelines
// @Summary 根据pipelineRun名称获取 taskRun
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data  true "根据pipelineRun名称获取 taskRun"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/getTaskRunList [get]
func GetTaskRunList(tr request.TektonTaskRunListUser, page, pageSize int) (err error, list interface{}, total int64){

	// create the clientset
	clientset, err := versioned.NewForConfig(global.GetK8sConfig(tr.UserName))
	if err != nil {
		panic(err.Error())
	}

	result, err := clientset.TektonV1alpha1().TaskRuns(tr.Namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: "tekton.dev/pipelineRun="+tr.PipelineRunName,
	})
	if err != nil {
		return
	}
	var taskRun model.TektonTaskRunList
	var taskStep model.TaskStep
	//var taskStepList []model.TaskStep
	listResult := make([]interface{}, 0)
	for _, t := range result.Items {
		taskRun.Name = t.Name
		taskRun.Namespace = t.Namespace
		taskRun.PodName = t.Status.PodName
		//taskRun.TaskRefName = t.Spec.TaskRef.Name
		taskRun.Status = string(t.Status.Conditions[0].Status)
		taskRun.Reason = t.Status.Conditions[0].Reason
		taskRun.Message = t.Status.Conditions[0].Message
		taskRun.TaskRefName = t.Labels["tekton.dev/pipelineTask"]
		taskRun.PipelineRunName = tr.PipelineRunName
		var taskStepList []model.TaskStep
		stepStatusFlag := false
		for _, step := range t.Status.Steps {
			taskStep.StepName = step.ContainerName
			taskStep.StepContainerName = step.ContainerName
			if step.Terminated != nil {
				taskStep.StartAt = step.Terminated.StartedAt
				taskStep.FinishAt = step.Terminated.FinishedAt
				taskStep.Reason = step.Terminated.Reason
			} else {
				if stepStatusFlag == false {
					taskStep.Reason = "Running"
					stepStatusFlag = true
				} else {
					taskStep.Reason = "Waiting"
				}
				taskStep.StartAt = step.Running.StartedAt
				taskStep.FinishAt = v1.Now()

			}

			taskStepList = append(taskStepList, taskStep)

		}
		taskRun.Steps = taskStepList
		if strings.Contains(t.Name, tr.Name) && strings.Contains(t.Namespace, tr.Namespace) {
			listResult = append(listResult, taskRun)
		}
	}
	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(page, pageSize, listResult)

	return err, list, total
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBuildPlan
//@description: 新建构建计划
//@param: api model.CreateBuildPlan
//@return: err error

func CreateBuildPlan(plan model.SysBuildPlan) (err error) {
	if !errors.Is(global.GVA_DB.Where("plan_name = ?", plan.PlanName).First(&model.SysBuildPlan{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同构建计划")
	}
	return global.GVA_DB.Create(&plan).Error
}


// GetBuildPlanList
// @Tags Pipelines
// @Summary 获取构建计划列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data  true "获取构建计划列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pipelines/getBuildPlanList [get]
func GetBuildPlanList(buildPlan model.SysBuildPlan, info request.PageInfo) (err error, list interface{}, total int64){

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysBuildPlan{})
	var buildPlanList []model.SysBuildPlan

	if buildPlan.PlanName != "" {
		db = db.Where("plan_name LIKE ?", "%"+buildPlan.PlanName+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, buildPlanList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("project_name").Find(&buildPlanList).Error
	}
	return err, buildPlanList, total
}


//@author: jun.huang
//@function: DeleteBuildPlan
//@description: 删除构建计划
//@param: api model.SysBuildPlan
//@return: err error

func DeleteBuildPlan(buildPlan model.SysBuildPlan) (err error) {
	//var user model.SysUser
	err = global.GVA_DB.Where("id = ?", buildPlan.ID).Delete(&buildPlan).Error
	return err
	//err = global.GVA_DB.Delete(&buildPlan).Error
	////ClearCasbin(1, api.Path, api.Method)
	//return err
}
