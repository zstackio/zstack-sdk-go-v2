// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2TicketFlow updates IAM2TicketFlow
func (cli *ZSClient) UpdateIAM2TicketFlow(uuid string, params param.UpdateIAM2TicketFlowParam) (*view.TicketFlowInventoryView, error) {
	resp := view.TicketFlowInventoryView{}
	if err := cli.Put("v1/tickets/flow", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2TicketFlow deletes IAM2TicketFlow
func (cli *ZSClient) DeleteIAM2TicketFlow(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow", uuid, string(deleteMode))
}
// AddIAM2TicketFlow adds IAM2TicketFlow
func (cli *ZSClient) AddIAM2TicketFlow(params param.AddIAM2TicketFlowParam) (*view.TicketFlowInventoryView, error) {
	resp := view.TicketFlowInventoryView{}
	if err := cli.Post("v1/tickets/flow", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
