// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeTicketStatus 操作TicketStatus
func (cli *ZSClient) ChangeTicketStatus(uuid string, params param.ChangeTicketStatusParam) (*view.ChangeTicketStatusEventView, error) {
	resp := view.ChangeTicketStatusEventView{}
	if err := cli.Put("v1/tickets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

