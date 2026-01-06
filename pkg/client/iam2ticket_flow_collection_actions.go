// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2TicketFlowCollection updates IAM2TicketFlowCollection
func (cli *ZSClient) UpdateIAM2TicketFlowCollection(uuid string, params param.UpdateIAM2TicketFlowCollectionParam) (*view.TicketFlowCollectionInventoryView, error) {
	var resp view.UpdateTicketFlowCollectionEventView
	if err := cli.Put("v1/tickets/flow-collections/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
