// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryExponBlockVolume queries ExponBlockVolume list
func (cli *ZSClient) QueryExponBlockVolume(params *param.QueryParam) ([]view.ExponBlockVolumeInventoryView, error) {
	var resp []view.ExponBlockVolumeInventoryView
	return resp, cli.List("v1/expon/block-volumes", params, &resp)
}

func (cli *ZSClient) GetExponBlockVolume(uuid string) (*view.ExponBlockVolumeInventoryView, error) {
	var resp view.ExponBlockVolumeInventoryView
	if err := cli.Get("v1/expon/block-volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
