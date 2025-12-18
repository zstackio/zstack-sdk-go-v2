// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2TickFlowCollection creates IAM2TickFlowCollection
func (cli *ZSClient) CreateIAM2TickFlowCollection(params param.CreateIAM2TickFlowCollectionParam) (*view.CreateTickFlowCollectionEventView, error) {
	resp := view.CreateTickFlowCollectionEventView{}
	if err := cli.Post("v1/tickets/flow-collections", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
