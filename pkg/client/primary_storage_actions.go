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

// PagePrimaryStorage Pagination
func (cli *ZSClient) PagePrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, int, error) {
	var primaryStorages []view.PrimaryStorageInventoryView
	total, err := cli.Page("v1/primary-storage", params, &primaryStorages)
	return primaryStorages, total, err
}
// ReconnectPrimaryStorage operates on PrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"reconnectPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(uuid string, params param.UpdatePrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"updatePrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePrimaryStorage deletes PrimaryStorage
func (cli *ZSClient) DeletePrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage", uuid, string(deleteMode))
}
