// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2ProjectRole creates IAM2ProjectRole
func (cli *ZSClient) CreateIAM2ProjectRole(params param.CreateIAM2ProjectRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.Post("v1/iam2/project-roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIAM2ProjectRole queries IAM2ProjectRole list
func (cli *ZSClient) QueryIAM2ProjectRole(params *param.QueryParam) ([]view.IAM2ProjectRoleInventoryView, error) {
	var resp []view.IAM2ProjectRoleInventoryView
	return resp, cli.List("v1/iam2/project-roles", params, &resp)
}

// PageIAM2ProjectRole Pagination
func (cli *ZSClient) PageIAM2ProjectRole(params *param.QueryParam) ([]view.IAM2ProjectRoleInventoryView, int, error) {
	var iAM2ProjectRoles []view.IAM2ProjectRoleInventoryView
	total, err := cli.Page("v1/iam2/project-roles", params, &iAM2ProjectRoles)
	return iAM2ProjectRoles, total, err
}
