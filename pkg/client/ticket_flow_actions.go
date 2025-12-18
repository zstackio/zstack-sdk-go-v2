// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTicketFlow 查询TicketFlow列表
func (cli *ZSClient) QueryTicketFlow(params param.QueryParam) ([]view.QueryTicketFlowView, error) {
	var resp []view.QueryTicketFlowView
	return resp, cli.List("v1/tickets/flow", &params, &resp)
}

