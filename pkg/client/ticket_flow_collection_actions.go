// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTicketFlowCollection queries TicketFlowCollection list
func (cli *ZSClient) QueryTicketFlowCollection(params *param.QueryParam) ([]view.TicketFlowCollectionInventoryView, error) {
	var resp []view.TicketFlowCollectionInventoryView
	return resp, cli.List("v1/tickets/flow-collections", params, &resp)
}
// DeleteTicketFlowCollection deletes TicketFlowCollection
func (cli *ZSClient) DeleteTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections/{uuid}", uuid, string(deleteMode))
}
