// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUser queries User list
func (cli *ZSClient) QueryUser(ctx context.Context, params *param.QueryParam) ([]view.UserInventoryView, error) {
	var resp []view.UserInventoryView
	return resp, cli.List(ctx, "v1/accounts/users", params, &resp)
}

func (cli *ZSClient) GetUser(ctx context.Context, uuid string) (*view.UserInventoryView, error) {
	var resp view.UserInventoryView
	if err := cli.Get(ctx, "v1/accounts/users", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUser Pagination
func (cli *ZSClient) PageUser(ctx context.Context, params *param.QueryParam) ([]view.UserInventoryView, int, error) {
	var users []view.UserInventoryView
	total, err := cli.Page(ctx, "v1/accounts/users", params, &users)
	return users, total, err
}
// DeleteUser deletes User
func (cli *ZSClient) DeleteUser(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts/users", uuid, string(deleteMode))
}
// UpdateUser updates User
func (cli *ZSClient) UpdateUser(ctx context.Context, params param.UpdateUserParam) (*view.UserInventoryView, error) {
	resp := view.UserInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/users/actions", "", "", map[string]interface{}{
		"updateUser": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateUser creates User
func (cli *ZSClient) CreateUser(ctx context.Context, params param.CreateUserParam) (*view.UserInventoryView, error) {
	resp := view.UserInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/accounts/users", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
