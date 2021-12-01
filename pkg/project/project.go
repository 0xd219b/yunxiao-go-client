package project

import (
	"fmt"

	"github.com/Berger7/yunxiao-go-client/config"
	rdc "github.com/aliyun/alibaba-cloud-sdk-go/services/devops-rdc"
)

const (
	defaultRegionID = "cn-hangzhou"
	defaultScheme   = "https"
	defaultPageSize = "10"
	defaultOrderBy  = "update_time"
)

const (
	ListAllProjects       = "all-projects"
	GetTaskIDsByProjectID = "all-tasks"
)

func RunProject() {
	switch config.Cfg.Action {
	case ListAllProjects:
		ListAllProjectsByOrgID(config.GetAK(), config.GetSK(), config.GetOrgID(), config.GetEndpoint())
	case GetTaskIDsByProjectID:
		GetAllTasksByProjectID(config.GetAK(), config.GetSK(), config.GetOrgID(), config.GetEndpoint(), config.GetProjectID())
	default:
		fmt.Println("no action")
	}
}

func GetAllTasksByProjectID(ak, sk, orgID, endpoint, projectID string) {
	client, err := rdc.NewClientWithAccessKey(defaultRegionID, ak, sk)
	if err != nil {
		panic(err)
	}

	req := rdc.CreateListDevopsProjectTaskListRequest()
	req.Scheme = defaultScheme

	req.ProjectId = projectID
	req.OrgId = orgID
	req.Domain = endpoint

	resp, err := client.ListDevopsProjectTaskList(req)
	if err != nil {
		panic(err)
	}
	if resp.Successful {
		listIterm := resp.Object.Result
		for _, item := range listIterm {
			fmt.Printf("%s\n", item.Id)
		}
	} else {
		fmt.Println(resp.ErrorMsg)
	}
}

func ListAllProjectsByOrgID(ak, sk, orgID, endpoint string) {
	client, err := rdc.NewClientWithAccessKey(defaultRegionID, ak, sk)
	if err != nil {
		panic(err)
	}

	req := rdc.CreateListDevopsProjectsRequest()
	req.Scheme = defaultScheme

	req.OrgId = orgID
	req.Domain = endpoint
	req.PageSize = defaultPageSize
	req.OrderBy = defaultOrderBy

	resp, err := client.ListDevopsProjects(req)
	if err != nil {
		panic(err)
	}
	if resp.Successful {
		listIterm := resp.Object.Result
		for _, item := range listIterm {
			fmt.Printf("%s %s\n", item.Id, item.Name)
		}
	} else {
		fmt.Println(resp.ErrorMsg)
	}
}
