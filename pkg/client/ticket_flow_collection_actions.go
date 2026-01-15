// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlowCollection queries TicketFlowCollection list
func (cli *ZSClient) QueryTicketFlowCollection(params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, error) {
	var resp []view.TicketFlowCollectionInventoryView
	return resp, cli.List("v1/tickets/flow-collections", params, &resp)
}

// PageTicketFlowCollection Pagination
func (cli *ZSClient) PageTicketFlowCollection(params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, int, error) {
	var ticketFlowCollections []view.TicketFlowCollectionInventoryView
	total, err := cli.Page("v1/tickets/flow-collections", params, &ticketFlowCollections)
	return ticketFlowCollections, total, err
}
// DeleteTicketFlowCollection deletes TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections", uuid, string(deleteMode))
}
