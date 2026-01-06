// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddExternalPrimaryStorage adds ExternalPrimaryStorage
func (cli *ZSClient) AddExternalPrimaryStorage(params param.AddExternalPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/addon", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateExternalPrimaryStorage updates ExternalPrimaryStorage
func (cli *ZSClient) UpdateExternalPrimaryStorage(uuid string, params param.UpdateExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	var resp view.UpdateExternalPrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/addon/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
