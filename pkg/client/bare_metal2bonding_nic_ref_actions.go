// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2BondingNicRef queries BareMetal2BondingNicRef list
func (cli *ZSClient) QueryBareMetal2BondingNicRef(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List("v1/baremetal2/bonding/nic/refs", params, &resp)
}

func (cli *ZSClient) GetBareMetal2BondingNicRef(uuid string) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.BareMetal2ChassisInventoryView
	if err := cli.Get("v1/baremetal2/bonding/nic/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
