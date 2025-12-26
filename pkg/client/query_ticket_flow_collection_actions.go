// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTicketFlowCollection queries TicketFlowCollection list
func (cli *ZSClient) QueryTicketFlowCollection(params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, error) {
	var resp []view.TicketFlowCollectionInventoryView
	return resp, cli.List("v1/tickets/flow-collections", params, &resp)
}
