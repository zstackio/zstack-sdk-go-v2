// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketType queries TicketType list
func (cli *ZSClient) QueryTicketType(ctx context.Context, params *param.QueryParam) ([]view.TicketTypeInventoryView, error) {
	var resp []view.TicketTypeInventoryView
	return resp, cli.List(ctx, "v1/ticket-types", params, &resp)
}

func (cli *ZSClient) GetTicketType(ctx context.Context, uuid string) (*view.TicketTypeInventoryView, error) {
	var resp view.TicketTypeInventoryView
	if err := cli.Get(ctx, "v1/ticket-types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicketType Pagination
func (cli *ZSClient) PageTicketType(ctx context.Context, params *param.QueryParam) ([]view.TicketTypeInventoryView, int, error) {
	var ticketTypes []view.TicketTypeInventoryView
	total, err := cli.Page(ctx, "v1/ticket-types", params, &ticketTypes)
	return ticketTypes, total, err
}
