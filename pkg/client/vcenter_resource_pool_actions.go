// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterResourcePool queries VCenterResourcePool list
func (cli *ZSClient) QueryVCenterResourcePool(params *param.QueryParam) ([]view.VCenterResourcePoolInventoryView, error) {
	var resp []view.VCenterResourcePoolInventoryView
	return resp, cli.List("v1/vcenters/clusters/resourcepools", params, &resp)
}

func (cli *ZSClient) GetVCenterResourcePool(uuid string) (*view.VCenterResourcePoolInventoryView, error) {
	var resp view.VCenterResourcePoolInventoryView
	if err := cli.Get("v1/vcenters/clusters/resourcepools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
