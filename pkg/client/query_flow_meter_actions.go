// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFlowMeter queries FlowMeter list
func (cli *ZSClient) QueryFlowMeter(params *param.QueryParam) ([]view.FlowMeterInventoryView, error) {
	var resp []view.FlowMeterInventoryView
	return resp, cli.List("v1/flowmeters", params, &resp)
}
