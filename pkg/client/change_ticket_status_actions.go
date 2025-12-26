// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeTicketStatus changes TicketStatus
func (cli *ZSClient) ChangeTicketStatus(uuid string, params param.ChangeTicketStatusParam) (*view.ChangeTicketStatusEventView, error) {
	resp := view.ChangeTicketStatusEventView{}
	if err := cli.Put("v1/tickets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
