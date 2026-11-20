// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2VirtualIDGroup queries IAM2VirtualIDGroup list
func (cli *ZSClient) QueryIAM2VirtualIDGroup(params *param.QueryParam) ([]view.IAM2VirtualIDGroupInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupInventoryView
	return resp, cli.List("v1/iam2/projects/groups", params, &resp)
}

func (cli *ZSClient) GetIAM2VirtualIDGroup(uuid string) (*view.IAM2VirtualIDGroupInventoryView, error) {
	var resp view.IAM2VirtualIDGroupInventoryView
	if err := cli.Get("v1/iam2/projects/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2VirtualIDGroup Pagination
func (cli *ZSClient) PageIAM2VirtualIDGroup(params *param.QueryParam) ([]view.IAM2VirtualIDGroupInventoryView, int, error) {
	var iAM2VirtualIDGroups []view.IAM2VirtualIDGroupInventoryView
	total, err := cli.Page("v1/iam2/projects/groups", params, &iAM2VirtualIDGroups)
	return iAM2VirtualIDGroups, total, err
}
// CreateIAM2VirtualIDGroup creates IAM2VirtualIDGroup
func (cli *ZSClient) CreateIAM2VirtualIDGroup(params param.CreateIAM2VirtualIDGroupParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	resp := view.IAM2VirtualIDGroupInventoryView{}
	if err := cli.Post("v1/iam2/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2VirtualIDGroup deletes IAM2VirtualIDGroup
func (cli *ZSClient) DeleteIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups", uuid, string(deleteMode))
}
// UpdateIAM2VirtualIDGroup updates IAM2VirtualIDGroup
func (cli *ZSClient) UpdateIAM2VirtualIDGroup(uuid string, params param.UpdateIAM2VirtualIDGroupParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	resp := view.IAM2VirtualIDGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/groups", uuid, "", map[string]interface{}{
		"updateIAM2VirtualIDGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDGroupToProjects adds IAM2VirtualIDGroupToProjects
func (cli *ZSClient) AddIAM2VirtualIDGroupToProjects(params param.AddIAM2VirtualIDGroupToProjectsParam) (*view.AddIAM2VirtualIDGroupToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDGroupToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveIAM2VirtualIDsFromGroup requires additional arguments.
//
// Deprecated: use RemoveIAM2VirtualIDsFromGroupWithParams to supply the required fields.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveIAM2VirtualIDsFromGroup requires additional arguments; use RemoveIAM2VirtualIDsFromGroupWithParams")
}

// RemoveIAM2VirtualIDsFromGroupWithParams sends the required route and query parameters for RemoveIAM2VirtualIDsFromGroup.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromGroupWithParams(groupUuid string, deleteMode param.DeleteMode, params param.RemoveIAM2VirtualIDsFromGroupParamDetail) error {
	query := url.Values{
		"virtualIDUuids": params.VirtualIDUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/projects/groups", groupUuid, "virtual-ids", deleteMode, query, "virtualIDUuids")
}

// RemoveIAM2VirtualIDGroupFromProjects requires additional arguments.
//
// Deprecated: use RemoveIAM2VirtualIDGroupFromProjectsWithParams to supply the required fields.
func (cli *ZSClient) RemoveIAM2VirtualIDGroupFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveIAM2VirtualIDGroupFromProjects requires additional arguments; use RemoveIAM2VirtualIDGroupFromProjectsWithParams")
}

// RemoveIAM2VirtualIDGroupFromProjectsWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveIAM2VirtualIDGroupFromProjectsWithParams(deleteMode param.DeleteMode, params param.RemoveIAM2VirtualIDGroupFromProjectsParamDetail) error {
	query := url.Values{
		"projectUuids": params.ProjectUuids,
		"groupUuids":   params.GroupUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/projects/groups", "", "", deleteMode, query, "projectUuids", "groupUuids")
}

// ChangeIAM2VirtualIDGroupState changes IAM2VirtualIDGroupState
func (cli *ZSClient) ChangeIAM2VirtualIDGroupState(uuid string, params param.ChangeIAM2VirtualIDGroupStateParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	resp := view.IAM2VirtualIDGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/groups", uuid, "", map[string]interface{}{
		"changeIAM2VirtualIDGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRolesToIAM2VirtualIDGroup adds RolesToIAM2VirtualIDGroup
func (cli *ZSClient) AddRolesToIAM2VirtualIDGroup(groupUuid string, params param.AddRolesToIAM2VirtualIDGroupParam) (*view.AddRolesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/projects/groups/%s/roles", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToGroup adds IAM2VirtualIDsToGroup
func (cli *ZSClient) AddIAM2VirtualIDsToGroup(groupUuid string, params param.AddIAM2VirtualIDsToGroupParam) (*view.AddIAM2VirtualIDToGroupEventView, error) {
	resp := view.AddIAM2VirtualIDToGroupEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/projects/groups/%s/virtual-ids", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRolesFromIAM2VirtualIDGroup requires additional arguments.
//
// Deprecated: use RemoveRolesFromIAM2VirtualIDGroupWithParams to supply the required fields.
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveRolesFromIAM2VirtualIDGroup requires additional arguments; use RemoveRolesFromIAM2VirtualIDGroupWithParams")
}

// RemoveRolesFromIAM2VirtualIDGroupWithParams sends the required route and query parameters for RemoveRolesFromIAM2VirtualIDGroup.
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDGroupWithParams(groupUuid string, deleteMode param.DeleteMode, params param.RemoveRolesFromIAM2VirtualIDGroupParamDetail) error {
	query := url.Values{
		"roleUuids": params.RoleUuids,
	}
	if params.ProjectUuid != nil {
		query.Set("projectUuid", *params.ProjectUuid)
	}
	return cli.deleteActionWithQuery("v1/iam2/projects/groups", groupUuid, "roles", deleteMode, query, "roleUuids")
}
