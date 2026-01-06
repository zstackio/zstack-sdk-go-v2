// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedBlockGroupPrimaryStorage queries SharedBlockGroupPrimaryStorage list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorage(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/sharedblockgroup", params, &resp)
}
// AddSharedBlockGroupPrimaryStorage adds SharedBlockGroupPrimaryStorage
func (cli *ZSClient) AddSharedBlockGroupPrimaryStorage(params param.AddSharedBlockGroupPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/sharedblockgroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
