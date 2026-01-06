// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2TicketFlow updates IAM2TicketFlow
func (cli *ZSClient) UpdateIAM2TicketFlow(uuid string, params param.UpdateIAM2TicketFlowParam) (*view.TicketFlowInventoryView, error) {
	var resp view.UpdateIAM2TicketFlowEventView
	if err := cli.Put("v1/tickets/flow/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIAM2TicketFlow deletes IAM2TicketFlow
func (cli *ZSClient) DeleteIAM2TicketFlow(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow/{uuid}", uuid, string(deleteMode))
}
// AddIAM2TicketFlow adds IAM2TicketFlow
func (cli *ZSClient) AddIAM2TicketFlow(params param.AddIAM2TicketFlowParam) (*view.TicketFlowInventoryView, error) {
	var resp view.AddIAM2TicketFlowEventView
	if err := cli.Post("v1/tickets/flow", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
