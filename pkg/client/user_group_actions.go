// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteUserGroup deletes UserGroup
func (cli *ZSClient) DeleteUserGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/accounts/groups", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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

func (cli *ZSClient) GetUserGroup(uuid string) (*view.UserGroupInventoryView, error) {
	var resp view.UserGroupInventoryView
	if err := cli.Get("v1/accounts/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(uuid string, params param.UpdateUserGroupParam) (*view.UserGroupInventoryView, error) {
	var resp view.UpdateUserGroupEventView
	if err := cli.Put("v1/accounts/groups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
