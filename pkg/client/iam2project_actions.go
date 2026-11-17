// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{}    // avoid unused import

// CreateIAM2Project creates IAM2Project
func (cli *ZSClient) CreateIAM2Project(params param.CreateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.Post("v1/iam2/projects", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIAM2Project deletes IAM2Project
func (cli *ZSClient) DeleteIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects", uuid, string(deleteMode))
}

// DeleteAndExpungeIAM2Project deletes IAM2Project and then expunges it from recycle bin.
func (cli *ZSClient) DeleteAndExpungeIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	if err := cli.DeleteIAM2Project(uuid, deleteMode); err != nil {
		return err
	}
	return cli.ExpungeIAM2Project(uuid)
}

// RecoverIAM2Project operates on IAM2Project
func (cli *ZSClient) RecoverIAM2Project(uuid string, params param.RecoverIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"recoverIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIAM2Project queries IAM2Project list
func (cli *ZSClient) QueryIAM2Project(params *param.QueryParam) ([]view.IAM2ProjectInventoryView, error) {
	var resp []view.IAM2ProjectInventoryView
	return resp, cli.List("v1/iam2/projects", params, &resp)
}

func (cli *ZSClient) GetIAM2Project(uuid string) (*view.IAM2ProjectInventoryView, error) {
	var resp view.IAM2ProjectInventoryView
	if err := cli.Get("v1/iam2/projects", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2Project Pagination
func (cli *ZSClient) PageIAM2Project(params *param.QueryParam) ([]view.IAM2ProjectInventoryView, int, error) {
	var iAM2Projects []view.IAM2ProjectInventoryView
	total, err := cli.Page("v1/iam2/projects", params, &iAM2Projects)
	return iAM2Projects, total, err
}

// ExpungeIAM2Project operates on IAM2Project
func (cli *ZSClient) ExpungeIAM2Project(uuid string) error {
	params := map[string]interface{}{
		"expungeIAM2Project": map[string]interface{}{},
	}
	return cli.Put("v1/iam2/projects", uuid, params, nil)
}

// LoginIAM2Project operates on IAM2Project
func (cli *ZSClient) LoginIAM2Project(params param.LoginIAM2ProjectParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/login", "", "", map[string]interface{}{
		"loginIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIAM2Project updates IAM2Project
func (cli *ZSClient) UpdateIAM2Project(uuid string, params param.UpdateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"updateIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopAllResourcesInIAM2Project stops AllResourcesInIAM2Project
func (cli *ZSClient) StopAllResourcesInIAM2Project(uuid string, params param.StopAllResourcesInIAM2ProjectParam) (*view.StopAllResourcesInIAM2ProjectEventView, error) {
	resp := view.StopAllResourcesInIAM2ProjectEventView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"stopAllResourcesInIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectContainerImageTags gets IAM2ProjectContainerImageTags by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImageTags(projectId string, repositoryId string, imageName string) (*view.ContainerImageTagInventoryView, error) {
	var resp view.ContainerImageTagInventoryView
	err := cli.GetWithSpec("v1/iam2/project", projectId, fmt.Sprintf("repository/%s/image/%s/tag", repositoryId, imageName), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectContainerClusterCandidates gets IAM2ProjectContainerClusterCandidates by uuid
func (cli *ZSClient) GetIAM2ProjectContainerClusterCandidates() (*view.ContainerClusterInventoryView, error) {
	var resp view.ContainerClusterInventoryView
	if err := cli.GetWithRespKey("v1/iam2/projects/container/cluster/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIAM2ProjectRetirePolicy operates on IAM2ProjectRetirePolicy
func (cli *ZSClient) SetIAM2ProjectRetirePolicy(uuid string, params param.SetIAM2ProjectRetirePolicyParam) (*view.SetIAM2ProjectRetirePolicyEventView, error) {
	resp := view.SetIAM2ProjectRetirePolicyEventView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/retire-policies", uuid, "", map[string]interface{}{
		"setIAM2ProjectRetirePolicy": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2ProjectFromTemplate creates IAM2ProjectFromTemplate
func (cli *ZSClient) CreateIAM2ProjectFromTemplate(templateUuid string, params param.CreateIAM2ProjectFromTemplateParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/projects/from/templates/%s", templateUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectRepository gets IAM2ProjectRepository by uuid
func (cli *ZSClient) GetIAM2ProjectRepository() (*view.ProjectRepositoryInventoryView, error) {
	var resp view.ProjectRepositoryInventoryView
	if err := cli.GetWithRespKey("v1/iam2/projects/repositories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIAM2ProjectContainerCluster operates on IAM2ProjectContainerCluster
func (cli *ZSClient) SetIAM2ProjectContainerCluster(uuid string, params param.SetIAM2ProjectContainerClusterParam) (*view.SetIAM2ProjectContainerClusterEventView, error) {
	resp := view.SetIAM2ProjectContainerClusterEventView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"setIAM2ProjectContainerCluster": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIAM2ProjectState changes IAM2ProjectState
func (cli *ZSClient) ChangeIAM2ProjectState(uuid string, params param.ChangeIAM2ProjectStateParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"changeIAM2ProjectState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectsOfVirtualID gets IAM2ProjectsOfVirtualID by uuid
func (cli *ZSClient) GetIAM2ProjectsOfVirtualID() (*view.IAM2ProjectInventoryView, error) {
	var resp view.IAM2ProjectInventoryView
	if err := cli.GetWithRespKey("v1/iam2/virtual-ids/projects", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourceToIAM2Project adds ResourceToIAM2Project
func (cli *ZSClient) AddResourceToIAM2Project(params param.AddResourceToIAM2ProjectParam) (*view.AddResourceToIAM2ProjectEventView, error) {
	resp := view.AddResourceToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/add/resource/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectContainerImages gets IAM2ProjectContainerImages by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImages(projectId string, repositoryId string) (*view.ZakuImageInventoryView, error) {
	var resp view.ZakuImageInventoryView
	err := cli.GetWithSpec("v1/iam2/project", projectId, fmt.Sprintf("repository/%s/image", repositoryId), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
