// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIAM2TicketFlow adds IAM2TicketFlow
func (cli *ZSClient) AddIAM2TicketFlow(params param.AddIAM2TicketFlowParam) (*view.AddIAM2TicketFlowEventView, error) {
	resp := view.AddIAM2TicketFlowEventView{}
	if err := cli.Post("v1/tickets/flow", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
