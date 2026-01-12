// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddBareMetal2IpmiChassis adds BareMetal2IpmiChassis
func (cli *ZSClient) AddBareMetal2IpmiChassis(params param.AddBareMetal2IpmiChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.AddBareMetal2ChassisEventView
	if err := cli.Post("v1/baremetal2/chassis/ipmi", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateBareMetal2IpmiChassis updates BareMetal2IpmiChassis
func (cli *ZSClient) UpdateBareMetal2IpmiChassis(uuid string, params param.UpdateBareMetal2IpmiChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.UpdateBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/ipmi", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
