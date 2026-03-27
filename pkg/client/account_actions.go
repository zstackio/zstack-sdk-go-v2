// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAccount updates Account
func (cli *ZSClient) UpdateAccount(ctx context.Context, uuid string, params param.UpdateAccountParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts", uuid, "", map[string]interface{}{
		"updateAccount": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccount deletes Account
func (cli *ZSClient) DeleteAccount(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts", uuid, string(deleteMode))
}
// CreateAccount creates Account
func (cli *ZSClient) CreateAccount(ctx context.Context, params param.CreateAccountParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.Post(ctx, "v1/accounts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAccount queries Account list
func (cli *ZSClient) QueryAccount(ctx context.Context, params *param.QueryParam) ([]view.AccountInventoryView, error) {
	var resp []view.AccountInventoryView
	return resp, cli.List(ctx, "v1/accounts", params, &resp)
}

func (cli *ZSClient) GetAccount(ctx context.Context, uuid string) (*view.AccountInventoryView, error) {
	var resp view.AccountInventoryView
	if err := cli.Get(ctx, "v1/accounts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccount Pagination
func (cli *ZSClient) PageAccount(ctx context.Context, params *param.QueryParam) ([]view.AccountInventoryView, int, error) {
	var accounts []view.AccountInventoryView
	total, err := cli.Page(ctx, "v1/accounts", params, &accounts)
	return accounts, total, err
}
