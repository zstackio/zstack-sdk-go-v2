// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeTicketFlowCollectionState changes TicketFlowCollectionState
func (cli *ZSClient) ChangeTicketFlowCollectionState(uuid string, params param.ChangeTicketFlowCollectionStateParam) (*view.ChangeTicketFlowCollectionStateEventView, error) {
	resp := view.ChangeTicketFlowCollectionStateEventView{}
	if err := cli.Put("v1/tickets/flow-collections/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
