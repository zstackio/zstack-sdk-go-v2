// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTicketHistory queries TicketHistory list
func (cli *ZSClient) QueryTicketHistory(params *param.QueryParam) ([]view.TicketStatusHistoryInventoryView, error) {
	var resp []view.TicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories", params, &resp)
}
