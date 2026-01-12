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
	var resp view.CreateTicketEventView
	if err := cli.Post("v1/tickets", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteTicket deletes Ticket
func (cli *ZSClient) DeleteTicket(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/tickets", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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
