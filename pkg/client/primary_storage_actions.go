// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPrimaryStorage queries PrimaryStorage list
func (cli *ZSClient) QueryPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage", params, &resp)
}

func (cli *ZSClient) GetPrimaryStorage(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectPrimaryStorage operates on PrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.ReconnectPrimaryStorageEventView
	err := cli.PutWithSpec("v1/primary-storage", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(uuid string, params param.UpdatePrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.UpdatePrimaryStorageEventView
	err := cli.PutWithSpec("v1/primary-storage", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePrimaryStorage deletes PrimaryStorage
func (cli *ZSClient) DeletePrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/primary-storage", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
