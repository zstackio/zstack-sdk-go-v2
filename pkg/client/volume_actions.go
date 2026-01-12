// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVolume updates Volume
func (cli *ZSClient) UpdateVolume(uuid string, params param.UpdateVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.UpdateVolumeEventView
	err := cli.PutWithSpec("v1/volumes", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVolume queries Volume list
func (cli *ZSClient) QueryVolume(params *param.QueryParam) ([]view.VolumeInventoryView, error) {
	var resp []view.VolumeInventoryView
	return resp, cli.List("v1/volumes", params, &resp)
}

func (cli *ZSClient) GetVolume(uuid string) (*view.VolumeInventoryView, error) {
	var resp view.VolumeInventoryView
	if err := cli.Get("v1/volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
