// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteUserGroup deletes UserGroup
func (cli *ZSClient) DeleteUserGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts/groups", uuid, string(deleteMode))
}
// CreateUserGroup creates UserGroup
func (cli *ZSClient) CreateUserGroup(ctx context.Context, params param.CreateUserGroupParam) (*view.UserGroupInventoryView, error) {
	resp := view.UserGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/accounts/groups", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryUserGroup queries UserGroup list
func (cli *ZSClient) QueryUserGroup(ctx context.Context, params *param.QueryParam) ([]view.UserGroupInventoryView, error) {
	var resp []view.UserGroupInventoryView
	return resp, cli.List(ctx, "v1/accounts/groups", params, &resp)
}

func (cli *ZSClient) GetUserGroup(ctx context.Context, uuid string) (*view.UserGroupInventoryView, error) {
	var resp view.UserGroupInventoryView
	if err := cli.Get(ctx, "v1/accounts/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUserGroup Pagination
func (cli *ZSClient) PageUserGroup(ctx context.Context, params *param.QueryParam) ([]view.UserGroupInventoryView, int, error) {
	var userGroups []view.UserGroupInventoryView
	total, err := cli.Page(ctx, "v1/accounts/groups", params, &userGroups)
	return userGroups, total, err
}
// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(ctx context.Context, params param.UpdateUserGroupParam) (*view.UserGroupInventoryView, error) {
	resp := view.UserGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/groups/actions", "", "", map[string]interface{}{
		"updateUserGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
