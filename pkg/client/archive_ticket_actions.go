// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryArchiveTicket 查询ArchiveTicket列表
func (cli *ZSClient) QueryArchiveTicket(params param.QueryParam) ([]view.QueryArchiveTicketView, error) {
	var resp []view.QueryArchiveTicketView
	return resp, cli.List("v1/tickets/archives", &params, &resp)
}

