// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTicketFlow queries TicketFlow list
func (cli *ZSClient) QueryTicketFlow(params *param.QueryParam) ([]view.TicketFlowInventoryView, error) {
	var resp []view.TicketFlowInventoryView
	return resp, cli.List("v1/tickets/flow", params, &resp)
}
