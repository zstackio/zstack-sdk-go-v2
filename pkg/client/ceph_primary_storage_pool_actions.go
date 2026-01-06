// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCephPrimaryStoragePool updates CephPrimaryStoragePool
func (cli *ZSClient) UpdateCephPrimaryStoragePool(uuid string, params param.UpdateCephPrimaryStoragePoolParam) (*view.CephPrimaryStoragePoolInventoryView, error) {
	var resp view.UpdateCephPrimaryStoragePoolEventView
	if err := cli.Put("v1/primary-storage/ceph/pools/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddCephPrimaryStoragePool adds CephPrimaryStoragePool
func (cli *ZSClient) AddCephPrimaryStoragePool(params param.AddCephPrimaryStoragePoolParam) (*view.CephPrimaryStoragePoolInventoryView, error) {
	var resp view.AddCephPrimaryStoragePoolEventView
	if err := cli.Post("v1/primary-storage/ceph/{primaryStorageUuid}/pools", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryCephPrimaryStoragePool queries CephPrimaryStoragePool list
func (cli *ZSClient) QueryCephPrimaryStoragePool(params *param.QueryParam) ([]view.CephPrimaryStoragePoolInventoryView, error) {
	var resp []view.CephPrimaryStoragePoolInventoryView
	return resp, cli.List("v1/primary-storage/ceph/pools", params, &resp)
}
// DeleteCephPrimaryStoragePool deletes CephPrimaryStoragePool
func (cli *ZSClient) DeleteCephPrimaryStoragePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/ceph/pools/{uuid}", uuid, string(deleteMode))
}
