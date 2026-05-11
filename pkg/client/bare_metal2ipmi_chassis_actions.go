// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddBareMetal2IpmiChassis adds BareMetal2IpmiChassis
func (cli *ZSClient) AddBareMetal2IpmiChassis(ctx context.Context, params param.AddBareMetal2IpmiChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal2/chassis/ipmi", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2IpmiChassis updates BareMetal2IpmiChassis
func (cli *ZSClient) UpdateBareMetal2IpmiChassis(ctx context.Context, uuid string, params param.UpdateBareMetal2IpmiChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis/ipmi", uuid, "", map[string]interface{}{
		"updateBareMetal2IpmiChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
