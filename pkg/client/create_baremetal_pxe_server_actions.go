// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBaremetalPxeServer creates BaremetalPxeServer
func (cli *ZSClient) CreateBaremetalPxeServer(params param.CreateBaremetalPxeServerParam) (*view.CreateBaremetalPxeServerEventView, error) {
	resp := view.CreateBaremetalPxeServerEventView{}
	if err := cli.Post("v1/baremetal/pxeservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
