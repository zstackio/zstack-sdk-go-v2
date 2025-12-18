// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFlowMeter creates FlowMeter
func (cli *ZSClient) CreateFlowMeter(params param.CreateFlowMeterParam) (*view.CreateFlowMeterEventView, error) {
	resp := view.CreateFlowMeterEventView{}
	if err := cli.Post("v1/flowmeters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
