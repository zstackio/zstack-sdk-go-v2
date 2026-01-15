// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAccount updates Account
func (cli *ZSClient) UpdateAccount(uuid string, params param.UpdateAccountParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.Put("v1/accounts", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccount deletes Account
func (cli *ZSClient) DeleteAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts", uuid, string(deleteMode))
}
// CreateAccount creates Account
func (cli *ZSClient) CreateAccount(params param.CreateAccountParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.Post("v1/accounts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAccount queries Account list
func (cli *ZSClient) QueryAccount(params *param.QueryParam) ([]view.AccountInventoryView, error) {
	var resp []view.AccountInventoryView
	return resp, cli.List("v1/accounts", params, &resp)
}

// PageAccount Pagination
func (cli *ZSClient) PageAccount(params *param.QueryParam) ([]view.AccountInventoryView, int, error) {
	var accounts []view.AccountInventoryView
	total, err := cli.Page("v1/accounts", params, &accounts)
	return accounts, total, err
}
