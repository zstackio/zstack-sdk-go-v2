// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIAM2TicketFlowCollection updates IAM2TicketFlowCollection
func (cli *ZSClient) UpdateIAM2TicketFlowCollection(uuid string, params param.UpdateIAM2TicketFlowCollectionParam) (*view.UpdateTicketFlowCollectionEventView, error) {
	resp := view.UpdateTicketFlowCollectionEventView{}
	if err := cli.Put("v1/tickets/flow-collections/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
