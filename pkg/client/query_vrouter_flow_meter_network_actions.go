// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVRouterFlowMeterNetwork queries VRouterFlowMeterNetwork list
func (cli *ZSClient) QueryVRouterFlowMeterNetwork(params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp []view.NetworkRouterFlowMeterRefInventoryView
	return resp, cli.List("v1/flowmeters/networks", params, &resp)
}
