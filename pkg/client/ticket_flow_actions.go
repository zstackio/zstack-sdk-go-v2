// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlow queries TicketFlow list
func (cli *ZSClient) QueryTicketFlow(ctx context.Context, params *param.QueryParam) ([]view.TicketFlowInventoryView, error) {
	var resp []view.TicketFlowInventoryView
	return resp, cli.List(ctx, "v1/tickets/flow", params, &resp)
}

func (cli *ZSClient) GetTicketFlow(ctx context.Context, uuid string) (*view.TicketFlowInventoryView, error) {
	var resp view.TicketFlowInventoryView
	if err := cli.Get(ctx, "v1/tickets/flow", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicketFlow Pagination
func (cli *ZSClient) PageTicketFlow(ctx context.Context, params *param.QueryParam) ([]view.TicketFlowInventoryView, int, error) {
	var ticketFlows []view.TicketFlowInventoryView
	total, err := cli.Page(ctx, "v1/tickets/flow", params, &ticketFlows)
	return ticketFlows, total, err
}
