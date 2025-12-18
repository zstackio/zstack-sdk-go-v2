// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateTicket 创建Ticket
func (cli *ZSClient) CreateTicket(params param.CreateTicketParam) (*view.CreateTicketEventView, error) {
	resp := view.CreateTicketEventView{}
	if err := cli.Post("v1/tickets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

