// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateTicketRequest 更新TicketRequest
func (cli *ZSClient) UpdateTicketRequest(uuid string, params param.UpdateTicketRequestParam) (*view.UpdateTicketRequestEventView, error) {
	resp := view.UpdateTicketRequestEventView{}
	if err := cli.Put("v1/tickets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

