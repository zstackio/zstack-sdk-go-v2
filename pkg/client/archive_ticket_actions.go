// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryArchiveTicket queries ArchiveTicket list
func (cli *ZSClient) QueryArchiveTicket(params *param.QueryParam) ([]view.ArchiveTicketInventoryView, error) {
	var resp []view.ArchiveTicketInventoryView
	return resp, cli.List("v1/tickets/archives", params, &resp)
}

// PageArchiveTicket Pagination
func (cli *ZSClient) PageArchiveTicket(params *param.QueryParam) ([]view.ArchiveTicketInventoryView, int, error) {
	var archiveTickets []view.ArchiveTicketInventoryView
	total, err := cli.Page("v1/tickets/archives", params, &archiveTickets)
	return archiveTickets, total, err
}
