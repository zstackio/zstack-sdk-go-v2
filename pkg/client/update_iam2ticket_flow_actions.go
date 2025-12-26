// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIAM2TicketFlow updates IAM2TicketFlow
func (cli *ZSClient) UpdateIAM2TicketFlow(uuid string, params param.UpdateIAM2TicketFlowParam) (*view.UpdateIAM2TicketFlowEventView, error) {
	resp := view.UpdateIAM2TicketFlowEventView{}
	if err := cli.Put("v1/tickets/flow/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
