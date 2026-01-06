// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
	return cli.Delete("v1/tickets/{uuid}", uuid, string(deleteMode))
}
// QueryTicket queries Ticket list
func (cli *ZSClient) QueryTicket(params *param.QueryParam) ([]view.TicketInventoryView, error) {
	var resp []view.TicketInventoryView
	return resp, cli.List("v1/tickets", params, &resp)
}
