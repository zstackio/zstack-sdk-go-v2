// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVRouterFlowMeterNetwork 查询VRouterFlowMeterNetwork列表
func (cli *ZSClient) QueryVRouterFlowMeterNetwork(params param.QueryParam) ([]view.QueryVRouterFlowMeterNetworkView, error) {
	var resp []view.QueryVRouterFlowMeterNetworkView
	return resp, cli.List("v1/flowmeters/networks", &params, &resp)
}

