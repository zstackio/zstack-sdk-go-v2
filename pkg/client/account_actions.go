// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAccount updates Account
func (cli *ZSClient) UpdateAccount(uuid string, params param.UpdateAccountParam) (*view.AccountInventoryView, error) {
	var resp view.UpdateAccountEventView
	if err := cli.Put("v1/accounts/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAccount deletes Account
func (cli *ZSClient) DeleteAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/{uuid}", uuid, string(deleteMode))
}
// CreateAccount creates Account
func (cli *ZSClient) CreateAccount(params param.CreateAccountParam) (*view.AccountInventoryView, error) {
	var resp view.CreateAccountEventView
	if err := cli.Post("v1/accounts", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAccount queries Account list
func (cli *ZSClient) QueryAccount(params *param.QueryParam) ([]view.AccountInventoryView, error) {
	var resp []view.AccountInventoryView
	return resp, cli.List("v1/accounts", params, &resp)
}
