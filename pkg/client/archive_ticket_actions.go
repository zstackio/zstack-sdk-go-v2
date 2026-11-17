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

func (cli *ZSClient) GetArchiveTicket(uuid string) (*view.ArchiveTicketInventoryView, error) {
	var resp view.ArchiveTicketInventoryView
	if err := cli.Get("v1/tickets/archives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageArchiveTicket Pagination
func (cli *ZSClient) PageArchiveTicket(params *param.QueryParam) ([]view.ArchiveTicketInventoryView, int, error) {
	var archiveTickets []view.ArchiveTicketInventoryView
	total, err := cli.Page("v1/tickets/archives", params, &archiveTickets)
	return archiveTickets, total, err
}

// QueryArchiveTicketHistory queries ArchiveTicketHistory list
func (cli *ZSClient) QueryArchiveTicketHistory(params *param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp []view.ArchiveTicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories/archives", params, &resp)
}

func (cli *ZSClient) GetArchiveTicketHistory(uuid string) (*view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp view.ArchiveTicketStatusHistoryInventoryView
	if err := cli.Get("v1/tickets/histories/archives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageArchiveTicketHistory Pagination
func (cli *ZSClient) PageArchiveTicketHistory(params *param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, int, error) {
	var archiveTicketHistories []view.ArchiveTicketStatusHistoryInventoryView
	total, err := cli.Page("v1/tickets/histories/archives", params, &archiveTicketHistories)
	return archiveTicketHistories, total, err
}
