// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2TicketFlowCollection 更新IAM2TicketFlowCollection
func (cli *ZSClient) UpdateIAM2TicketFlowCollection(uuid string, params param.UpdateIAM2TicketFlowCollectionParam) (*view.UpdateTicketFlowCollectionEventView, error) {
	resp := view.UpdateTicketFlowCollectionEventView{}
	if err := cli.Put("v1/tickets/flow-collections/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

