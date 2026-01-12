// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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
	err := cli.PutWithSpec("v1/primary-storage/addon", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
