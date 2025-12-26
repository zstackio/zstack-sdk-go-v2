// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectBaremetalPxeServer operates on ReconnectBaremetalPxeServer
func (cli *ZSClient) ReconnectBaremetalPxeServer(uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.ReconnectBaremetalPxeServerEventView, error) {
	resp := view.ReconnectBaremetalPxeServerEventView{}
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
