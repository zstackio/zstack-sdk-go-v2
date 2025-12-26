// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFlowMeter creates FlowMeter
func (cli *ZSClient) CreateFlowMeter(params param.CreateFlowMeterParam) (*view.CreateFlowMeterEventView, error) {
	resp := view.CreateFlowMeterEventView{}
	if err := cli.Post("v1/flowmeters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
