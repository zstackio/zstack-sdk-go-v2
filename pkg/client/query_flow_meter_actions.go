// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFlowMeter queries FlowMeter list
func (cli *ZSClient) QueryFlowMeter(params param.QueryParam) ([]view.FlowMeterInventoryView, error) {
	var resp []view.FlowMeterInventoryView
	return resp, cli.List("v1/flowmeters", &params, &resp)
}
