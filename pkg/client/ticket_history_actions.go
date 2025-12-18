// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTicketHistory 查询TicketHistory列表
func (cli *ZSClient) QueryTicketHistory(params param.QueryParam) ([]view.QueryTicketHistoryView, error) {
	var resp []view.QueryTicketHistoryView
	return resp, cli.List("v1/tickets/histories", &params, &resp)
}

