// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTicketType 查询TicketType列表
func (cli *ZSClient) QueryTicketType(params param.QueryParam) ([]view.QueryTicketTypeView, error) {
	var resp []view.QueryTicketTypeView
	return resp, cli.List("v1/ticket-types", &params, &resp)
}

