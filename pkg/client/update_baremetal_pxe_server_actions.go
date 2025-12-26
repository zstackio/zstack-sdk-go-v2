// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBaremetalPxeServer updates BaremetalPxeServer
func (cli *ZSClient) UpdateBaremetalPxeServer(uuid string, params param.UpdateBaremetalPxeServerParam) (*view.UpdateBaremetalPxeServerEventView, error) {
	resp := view.UpdateBaremetalPxeServerEventView{}
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
