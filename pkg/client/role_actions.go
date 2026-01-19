// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateRole creates Role
func (cli *ZSClient) CreateRole(params param.CreateRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.Post("v1/identities/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryRole queries Role list
func (cli *ZSClient) QueryRole(params *param.QueryParam) ([]view.RoleInventoryView, error) {
	var resp []view.RoleInventoryView
	return resp, cli.List("v1/identities/roles", params, &resp)
}

func (cli *ZSClient) GetRole(uuid string) (*view.RoleInventoryView, error) {
	var resp view.RoleInventoryView
	if err := cli.Get("v1/identities/roles", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageRole Pagination
func (cli *ZSClient) PageRole(params *param.QueryParam) ([]view.RoleInventoryView, int, error) {
	var roles []view.RoleInventoryView
	total, err := cli.Page("v1/identities/roles", params, &roles)
	return roles, total, err
}
// DeleteRole deletes Role
func (cli *ZSClient) DeleteRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles", uuid, string(deleteMode))
}
// UpdateRole updates Role
func (cli *ZSClient) UpdateRole(uuid string, params param.UpdateRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.Put("v1/identities/roles", uuid, map[string]interface{}{
		"updateRole": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
