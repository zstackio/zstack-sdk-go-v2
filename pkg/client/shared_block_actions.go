// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedBlock queries SharedBlock list
func (cli *ZSClient) QuerySharedBlock(params *param.QueryParam) ([]view.SharedBlockInventoryView, error) {
	var resp []view.SharedBlockInventoryView
	return resp, cli.List("v1/sharedblock-group/sharedblocks", params, &resp)
}

func (cli *ZSClient) GetSharedBlock(uuid string) (*view.SharedBlockInventoryView, error) {
	var resp view.SharedBlockInventoryView
	if err := cli.Get("v1/sharedblock-group/sharedblocks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSharedBlock updates SharedBlock
func (cli *ZSClient) UpdateSharedBlock(uuid string, params param.UpdateSharedBlockParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp view.UpdateSharedBlockEventView
	if err := cli.Put("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
