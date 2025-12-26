// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBaremetalPxeServer creates BaremetalPxeServer
func (cli *ZSClient) CreateBaremetalPxeServer(params param.CreateBaremetalPxeServerParam) (*view.CreateBaremetalPxeServerEventView, error) {
	resp := view.CreateBaremetalPxeServerEventView{}
	if err := cli.Post("v1/baremetal/pxeservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
