// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2TicketFlowCollection updates IAM2TicketFlowCollection
func (cli *ZSClient) UpdateIAM2TicketFlowCollection(uuid string, params param.UpdateIAM2TicketFlowCollectionParam) (*view.TicketFlowCollectionInventoryView, error) {
	resp := view.TicketFlowCollectionInventoryView{}
	if err := cli.PutWithRespKey("v1/tickets/flow-collections", uuid, "", map[string]interface{}{
		"updateIAM2TicketFlowCollection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
