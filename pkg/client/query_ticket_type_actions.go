// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTicketType queries TicketType list
func (cli *ZSClient) QueryTicketType(params *param.QueryParam) ([]view.TicketTypeInventoryView, error) {
	var resp []view.TicketTypeInventoryView
	return resp, cli.List("v1/ticket-types", params, &resp)
}
