// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVRouterNetworksToFlowMeter adds VRouterNetworksToFlowMeter
func (cli *ZSClient) AddVRouterNetworksToFlowMeter(params param.AddVRouterNetworksToFlowMeterParam) (*view.AddVRouterNetworksToFlowMeterEventView, error) {
	resp := view.AddVRouterNetworksToFlowMeterEventView{}
	if err := cli.Post("v1/flowmeters/{flowMeterUuid}/router/{vRouterUuid}/addnetworks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
