// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryArchiveTicketHistory 查询ArchiveTicketHistory列表
func (cli *ZSClient) QueryArchiveTicketHistory(params param.QueryParam) ([]view.QueryArchiveTicketHistoryView, error) {
	var resp []view.QueryArchiveTicketHistoryView
	return resp, cli.List("v1/tickets/histories/archives", &params, &resp)
}

