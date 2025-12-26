// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateTicket creates Ticket
func (cli *ZSClient) CreateTicket(params param.CreateTicketParam) (*view.CreateTicketEventView, error) {
	resp := view.CreateTicketEventView{}
	if err := cli.Post("v1/tickets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
