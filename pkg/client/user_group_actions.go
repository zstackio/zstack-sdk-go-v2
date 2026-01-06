// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteUserGroup deletes UserGroup
func (cli *ZSClient) DeleteUserGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{uuid}", uuid, string(deleteMode))
}
// CreateUserGroup creates UserGroup
func (cli *ZSClient) CreateUserGroup(params param.CreateUserGroupParam) (*view.UserGroupInventoryView, error) {
	var resp view.CreateUserGroupEventView
	if err := cli.Post("v1/accounts/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryUserGroup queries UserGroup list
func (cli *ZSClient) QueryUserGroup(params *param.QueryParam) ([]view.UserGroupInventoryView, error) {
	var resp []view.UserGroupInventoryView
	return resp, cli.List("v1/accounts/groups", params, &resp)
}
// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(uuid string, params param.UpdateUserGroupParam) (*view.UserGroupInventoryView, error) {
	var resp view.UpdateUserGroupEventView
	if err := cli.Put("v1/accounts/groups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
