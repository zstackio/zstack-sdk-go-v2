// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCephPrimaryStoragePool updates CephPrimaryStoragePool
func (cli *ZSClient) UpdateCephPrimaryStoragePool(uuid string, params param.UpdateCephPrimaryStoragePoolParam) (*view.CephPrimaryStoragePoolInventoryView, error) {
	resp := view.CephPrimaryStoragePoolInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage/ceph/pools", uuid, "", map[string]interface{}{
		"updateCephPrimaryStoragePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddCephPrimaryStoragePool adds CephPrimaryStoragePool
func (cli *ZSClient) AddCephPrimaryStoragePool(primaryStorageUuid string, params param.AddCephPrimaryStoragePoolParam) (*view.CephPrimaryStoragePoolInventoryView, error) {
	resp := view.CephPrimaryStoragePoolInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/primary-storage/ceph/%s/pools", primaryStorageUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCephPrimaryStoragePool queries CephPrimaryStoragePool list
func (cli *ZSClient) QueryCephPrimaryStoragePool(params *param.QueryParam) ([]view.CephPrimaryStoragePoolInventoryView, error) {
	var resp []view.CephPrimaryStoragePoolInventoryView
	return resp, cli.List("v1/primary-storage/ceph/pools", params, &resp)
}

func (cli *ZSClient) GetCephPrimaryStoragePool(uuid string) (*view.CephPrimaryStoragePoolInventoryView, error) {
	var resp view.CephPrimaryStoragePoolInventoryView
	if err := cli.Get("v1/primary-storage/ceph/pools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephPrimaryStoragePool Pagination
func (cli *ZSClient) PageCephPrimaryStoragePool(params *param.QueryParam) ([]view.CephPrimaryStoragePoolInventoryView, int, error) {
	var cephPrimaryStoragePools []view.CephPrimaryStoragePoolInventoryView
	total, err := cli.Page("v1/primary-storage/ceph/pools", params, &cephPrimaryStoragePools)
	return cephPrimaryStoragePools, total, err
}
// DeleteCephPrimaryStoragePool deletes CephPrimaryStoragePool
func (cli *ZSClient) DeleteCephPrimaryStoragePool(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/ceph/pools", uuid, string(deleteMode))
}
