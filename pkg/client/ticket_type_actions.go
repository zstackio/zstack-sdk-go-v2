// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketType queries TicketType list
func (cli *ZSClient) QueryTicketType(params *param.QueryParam) ([]view.TicketTypeInventoryView, error) {
	var resp []view.TicketTypeInventoryView
	return resp, cli.List("v1/ticket-types", params, &resp)
}

// PageTicketType Pagination
func (cli *ZSClient) PageTicketType(params *param.QueryParam) ([]view.TicketTypeInventoryView, int, error) {
	var ticketTypes []view.TicketTypeInventoryView
	total, err := cli.Page("v1/ticket-types", params, &ticketTypes)
	return ticketTypes, total, err
}
