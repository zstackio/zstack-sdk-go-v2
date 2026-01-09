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
	var resp view.CreateRoleEventView
	if err := cli.Post("v1/identities/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
// DeleteRole deletes Role
func (cli *ZSClient) DeleteRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles", uuid, string(deleteMode))
}
// UpdateRole updates Role
func (cli *ZSClient) UpdateRole(uuid string, params param.UpdateRoleParam) (*view.RoleInventoryView, error) {
	var resp view.UpdateRoleEventView
	if err := cli.Put("v1/identities/roles/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
