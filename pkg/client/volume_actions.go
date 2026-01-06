// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVolume updates Volume
func (cli *ZSClient) UpdateVolume(uuid string, params param.UpdateVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.UpdateVolumeEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVolume queries Volume list
func (cli *ZSClient) QueryVolume(params *param.QueryParam) ([]view.VolumeInventoryView, error) {
	var resp []view.VolumeInventoryView
	return resp, cli.List("v1/volumes", params, &resp)
}
