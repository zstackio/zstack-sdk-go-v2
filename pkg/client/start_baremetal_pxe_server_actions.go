// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StartBaremetalPxeServer starts BaremetalPxeServer
func (cli *ZSClient) StartBaremetalPxeServer(uuid string, params param.StartBaremetalPxeServerParam) (*view.StartBaremetalPxeServerEventView, error) {
	resp := view.StartBaremetalPxeServerEventView{}
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
