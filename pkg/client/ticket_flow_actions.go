// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlow queries TicketFlow list
func (cli *ZSClient) QueryTicketFlow(params *param.QueryParam) ([]view.TicketFlowInventoryView, error) {
	var resp []view.TicketFlowInventoryView
	return resp, cli.List("v1/tickets/flow", params, &resp)
}

// PageTicketFlow Pagination
func (cli *ZSClient) PageTicketFlow(params *param.QueryParam) ([]view.TicketFlowInventoryView, int, error) {
	var ticketFlows []view.TicketFlowInventoryView
	total, err := cli.Page("v1/tickets/flow", params, &ticketFlows)
	return ticketFlows, total, err
}
