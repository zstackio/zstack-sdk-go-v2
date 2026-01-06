// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAccessControlList creates AccessControlList
func (cli *ZSClient) CreateAccessControlList(params param.CreateAccessControlListParam) (*view.AccessControlListInventoryView, error) {
	var resp view.CreateAccessControlListEventView
	if err := cli.Post("v1/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAccessControlList updates AccessControlList
func (cli *ZSClient) UpdateAccessControlList(uuid string, params param.UpdateAccessControlListParam) (*view.AccessControlListInventoryView, error) {
	var resp view.UpdateAccessControlListEventView
	if err := cli.Put("v1/access-control-lists/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAccessControlList queries AccessControlList list
func (cli *ZSClient) QueryAccessControlList(params *param.QueryParam) ([]view.AccessControlListInventoryView, error) {
	var resp []view.AccessControlListInventoryView
	return resp, cli.List("v1/access-control-lists", params, &resp)
}
// DeleteAccessControlList deletes AccessControlList
func (cli *ZSClient) DeleteAccessControlList(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/access-control-lists/{uuid}", uuid, string(deleteMode))
}
