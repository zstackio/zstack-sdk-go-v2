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
	return cli.Delete("v1/accounts/groups", uuid, string(deleteMode))
}
// CreateUserGroup creates UserGroup
func (cli *ZSClient) CreateUserGroup(params param.CreateUserGroupParam) (*view.UserGroupInventoryView, error) {
	resp := view.UserGroupInventoryView{}
	if err := cli.Post("v1/accounts/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryUserGroup queries UserGroup list
func (cli *ZSClient) QueryUserGroup(params *param.QueryParam) ([]view.UserGroupInventoryView, error) {
	var resp []view.UserGroupInventoryView
	return resp, cli.List("v1/accounts/groups", params, &resp)
}

// PageUserGroup Pagination
func (cli *ZSClient) PageUserGroup(params *param.QueryParam) ([]view.UserGroupInventoryView, int, error) {
	var userGroups []view.UserGroupInventoryView
	total, err := cli.Page("v1/accounts/groups", params, &userGroups)
	return userGroups, total, err
}
// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(uuid string, params param.UpdateUserGroupParam) (*view.UserGroupInventoryView, error) {
	resp := view.UserGroupInventoryView{}
	if err := cli.Put("v1/accounts/groups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
