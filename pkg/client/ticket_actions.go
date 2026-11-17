// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateTicket creates Ticket
func (cli *ZSClient) CreateTicket(params param.CreateTicketParam) (*view.TicketInventoryView, error) {
	resp := view.TicketInventoryView{}
	if err := cli.Post("v1/tickets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteTicket deletes Ticket
func (cli *ZSClient) DeleteTicket(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets", uuid, string(deleteMode))
}
// QueryTicket queries Ticket list
func (cli *ZSClient) QueryTicket(params *param.QueryParam) ([]view.TicketInventoryView, error) {
	var resp []view.TicketInventoryView
	return resp, cli.List("v1/tickets", params, &resp)
}

func (cli *ZSClient) GetTicket(uuid string) (*view.TicketInventoryView, error) {
	var resp view.TicketInventoryView
	if err := cli.Get("v1/tickets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicket Pagination
func (cli *ZSClient) PageTicket(params *param.QueryParam) ([]view.TicketInventoryView, int, error) {
	var tickets []view.TicketInventoryView
	total, err := cli.Page("v1/tickets", params, &tickets)
	return tickets, total, err
}

// QueryTicketHistory queries TicketHistory list
func (cli *ZSClient) QueryTicketHistory(params *param.QueryParam) ([]view.TicketStatusHistoryInventoryView, error) {
	var resp []view.TicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories", params, &resp)
}

func (cli *ZSClient) GetTicketHistory(uuid string) (*view.TicketStatusHistoryInventoryView, error) {
	var resp view.TicketStatusHistoryInventoryView
	if err := cli.Get("v1/tickets/histories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicketHistory Pagination
func (cli *ZSClient) PageTicketHistory(params *param.QueryParam) ([]view.TicketStatusHistoryInventoryView, int, error) {
	var ticketHistories []view.TicketStatusHistoryInventoryView
	total, err := cli.Page("v1/tickets/histories", params, &ticketHistories)
	return ticketHistories, total, err
}

// UpdateTicketRequest updates TicketRequest
func (cli *ZSClient) UpdateTicketRequest(uuid string, params param.UpdateTicketRequestParam) (*view.TicketInventoryView, error) {
	resp := view.TicketInventoryView{}
	if err := cli.PutWithRespKey("v1/tickets", uuid, "", map[string]interface{}{
		"updateTicketRequest": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeTicketStatus changes TicketStatus
func (cli *ZSClient) ChangeTicketStatus(uuid string, params param.ChangeTicketStatusParam) (*view.TicketInventoryView, error) {
	resp := view.TicketInventoryView{}
	if err := cli.PutWithRespKey("v1/tickets", uuid, "", map[string]interface{}{
		"changeTicketStatus": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
