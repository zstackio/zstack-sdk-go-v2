// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIAM2TicketFlow 操作AddIAM2TicketFlow
func (cli *ZSClient) AddIAM2TicketFlow(params param.AddIAM2TicketFlowParam) (*view.AddIAM2TicketFlowEventView, error) {
	resp := view.AddIAM2TicketFlowEventView{}
	if err := cli.Post("v1/tickets/flow", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

