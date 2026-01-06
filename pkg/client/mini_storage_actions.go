// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMiniStorage queries MiniStorage list
func (cli *ZSClient) QueryMiniStorage(params *param.QueryParam) ([]view.MiniStorageInventoryView, error) {
	var resp []view.MiniStorageInventoryView
	return resp, cli.List("v1/primary-storage/mini", params, &resp)
}
// AddMiniStorage adds MiniStorage
func (cli *ZSClient) AddMiniStorage(params param.AddMiniStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/mini", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
