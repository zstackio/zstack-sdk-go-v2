// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateTicketRequest updates TicketRequest
func (cli *ZSClient) UpdateTicketRequest(uuid string, params param.UpdateTicketRequestParam) (*view.UpdateTicketRequestEventView, error) {
	resp := view.UpdateTicketRequestEventView{}
	if err := cli.Put("v1/tickets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
