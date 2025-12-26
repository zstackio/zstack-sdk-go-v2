// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTicket queries Ticket list
func (cli *ZSClient) QueryTicket(params *param.QueryParam) ([]view.TicketInventoryView, error) {
	var resp []view.TicketInventoryView
	return resp, cli.List("v1/tickets", params, &resp)
}
