// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryArchiveTicket queries ArchiveTicket list
func (cli *ZSClient) QueryArchiveTicket(ctx context.Context, params *param.QueryParam) ([]view.ArchiveTicketInventoryView, error) {
	var resp []view.ArchiveTicketInventoryView
	return resp, cli.List(ctx, "v1/tickets/archives", params, &resp)
}

func (cli *ZSClient) GetArchiveTicket(ctx context.Context, uuid string) (*view.ArchiveTicketInventoryView, error) {
	var resp view.ArchiveTicketInventoryView
	if err := cli.Get(ctx, "v1/tickets/archives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageArchiveTicket Pagination
func (cli *ZSClient) PageArchiveTicket(ctx context.Context, params *param.QueryParam) ([]view.ArchiveTicketInventoryView, int, error) {
	var archiveTickets []view.ArchiveTicketInventoryView
	total, err := cli.Page(ctx, "v1/tickets/archives", params, &archiveTickets)
	return archiveTickets, total, err
}
