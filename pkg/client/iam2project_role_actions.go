// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2ProjectRole creates IAM2ProjectRole
func (cli *ZSClient) CreateIAM2ProjectRole(params param.CreateIAM2ProjectRoleParam) (*view.RoleInventoryView, error) {
	var resp view.CreateRoleEventView
	if err := cli.Post("v1/iam2/project-roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIAM2ProjectRole queries IAM2ProjectRole list
func (cli *ZSClient) QueryIAM2ProjectRole(params *param.QueryParam) ([]view.IAM2ProjectRoleInventoryView, error) {
	var resp []view.IAM2ProjectRoleInventoryView
	return resp, cli.List("v1/iam2/project-roles", params, &resp)
}
