// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2Gateway 查询BareMetal2Gateway列表
func (cli *ZSClient) QueryBareMetal2Gateway(params param.QueryParam) ([]view.QueryBareMetal2GatewayView, error) {
	var resp []view.QueryBareMetal2GatewayView
	return resp, cli.List("v1/baremetal2/gateways", &params, &resp)
}

