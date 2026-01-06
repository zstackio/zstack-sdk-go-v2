// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedBlock queries SharedBlock list
func (cli *ZSClient) QuerySharedBlock(params *param.QueryParam) ([]view.SharedBlockInventoryView, error) {
	var resp []view.SharedBlockInventoryView
	return resp, cli.List("v1/sharedblock-group/sharedblocks", params, &resp)
}
// UpdateSharedBlock updates SharedBlock
func (cli *ZSClient) UpdateSharedBlock(uuid string, params param.UpdateSharedBlockParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp view.UpdateSharedBlockEventView
	if err := cli.Put("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
