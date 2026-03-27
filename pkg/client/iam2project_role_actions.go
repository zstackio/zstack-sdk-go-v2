// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2ProjectRole creates IAM2ProjectRole
func (cli *ZSClient) CreateIAM2ProjectRole(ctx context.Context) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/project-roles", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIAM2ProjectRole queries IAM2ProjectRole list
func (cli *ZSClient) QueryIAM2ProjectRole(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectRoleInventoryView, error) {
	var resp []view.IAM2ProjectRoleInventoryView
	return resp, cli.List(ctx, "v1/iam2/project-roles", params, &resp)
}

func (cli *ZSClient) GetIAM2ProjectRole(ctx context.Context, uuid string) (*view.IAM2ProjectRoleInventoryView, error) {
	var resp view.IAM2ProjectRoleInventoryView
	if err := cli.Get(ctx, "v1/iam2/project-roles", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2ProjectRole Pagination
func (cli *ZSClient) PageIAM2ProjectRole(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectRoleInventoryView, int, error) {
	var iAM2ProjectRoles []view.IAM2ProjectRoleInventoryView
	total, err := cli.Page(ctx, "v1/iam2/project-roles", params, &iAM2ProjectRoles)
	return iAM2ProjectRoles, total, err
}
