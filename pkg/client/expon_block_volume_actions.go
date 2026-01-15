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

// PageExponBlockVolume Pagination
func (cli *ZSClient) PageExponBlockVolume(params *param.QueryParam) ([]view.ExponBlockVolumeInventoryView, int, error) {
	var exponBlockVolumes []view.ExponBlockVolumeInventoryView
	total, err := cli.Page("v1/expon/block-volumes", params, &exponBlockVolumes)
	return exponBlockVolumes, total, err
}
