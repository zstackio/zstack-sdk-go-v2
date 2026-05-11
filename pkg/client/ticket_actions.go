// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateTicket creates Ticket
func (cli *ZSClient) CreateTicket(ctx context.Context, params param.CreateTicketParam) (*view.TicketInventoryView, error) {
	resp := view.TicketInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/tickets", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteTicket deletes Ticket
func (cli *ZSClient) DeleteTicket(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tickets", uuid, string(deleteMode))
}
// QueryTicket queries Ticket list
func (cli *ZSClient) QueryTicket(ctx context.Context, params *param.QueryParam) ([]view.TicketInventoryView, error) {
	var resp []view.TicketInventoryView
	return resp, cli.List(ctx, "v1/tickets", params, &resp)
}

func (cli *ZSClient) GetTicket(ctx context.Context, uuid string) (*view.TicketInventoryView, error) {
	var resp view.TicketInventoryView
	if err := cli.Get(ctx, "v1/tickets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicket Pagination
func (cli *ZSClient) PageTicket(ctx context.Context, params *param.QueryParam) ([]view.TicketInventoryView, int, error) {
	var tickets []view.TicketInventoryView
	total, err := cli.Page(ctx, "v1/tickets", params, &tickets)
	return tickets, total, err
}
