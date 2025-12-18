// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2TicketFlow 更新IAM2TicketFlow
func (cli *ZSClient) UpdateIAM2TicketFlow(uuid string, params param.UpdateIAM2TicketFlowParam) (*view.UpdateIAM2TicketFlowEventView, error) {
	resp := view.UpdateIAM2TicketFlowEventView{}
	if err := cli.Put("v1/tickets/flow/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

