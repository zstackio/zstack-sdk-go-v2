// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPrimaryStorage queries PrimaryStorage list
func (cli *ZSClient) QueryPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage", params, &resp)
}
// ReconnectPrimaryStorage operates on PrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.ReconnectPrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(uuid string, params param.UpdatePrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.UpdatePrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePrimaryStorage deletes PrimaryStorage
func (cli *ZSClient) DeletePrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/{uuid}", uuid, string(deleteMode))
}
