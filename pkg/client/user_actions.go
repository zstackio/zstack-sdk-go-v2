// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUser queries User list
func (cli *ZSClient) QueryUser(params *param.QueryParam) ([]view.UserInventoryView, error) {
	var resp []view.UserInventoryView
	return resp, cli.List("v1/accounts/users", params, &resp)
}

func (cli *ZSClient) GetUser(uuid string) (*view.UserInventoryView, error) {
	var resp view.UserInventoryView
	if err := cli.Get("v1/accounts/users", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteUser deletes User
func (cli *ZSClient) DeleteUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/accounts/users", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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
