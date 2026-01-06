// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUser queries User list
func (cli *ZSClient) QueryUser(params *param.QueryParam) ([]view.UserInventoryView, error) {
	var resp []view.UserInventoryView
	return resp, cli.List("v1/accounts/users", params, &resp)
}
// DeleteUser deletes User
func (cli *ZSClient) DeleteUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{uuid}", uuid, string(deleteMode))
}
// UpdateUser updates User
func (cli *ZSClient) UpdateUser(uuid string, params param.UpdateUserParam) (*view.UserInventoryView, error) {
	var resp view.UpdateUserEventView
	if err := cli.Put("v1/accounts/users/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateUser creates User
func (cli *ZSClient) CreateUser(params param.CreateUserParam) (*view.UserInventoryView, error) {
	var resp view.CreateUserEventView
	if err := cli.Post("v1/accounts/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
