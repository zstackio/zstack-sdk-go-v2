// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddCephPrimaryStorage adds CephPrimaryStorage
func (cli *ZSClient) AddCephPrimaryStorage(params param.AddCephPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryCephPrimaryStorage queries CephPrimaryStorage list
func (cli *ZSClient) QueryCephPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/ceph", params, &resp)
}
