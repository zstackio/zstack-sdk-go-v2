// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StopBaremetalPxeServer stops BaremetalPxeServer
func (cli *ZSClient) StopBaremetalPxeServer(uuid string, params param.StopBaremetalPxeServerParam) (*view.StopBaremetalPxeServerEventView, error) {
	resp := view.StopBaremetalPxeServerEventView{}
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
