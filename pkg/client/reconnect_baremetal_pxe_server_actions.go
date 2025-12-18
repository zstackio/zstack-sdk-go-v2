// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectBaremetalPxeServer 操作ReconnectBaremetalPxeServer
func (cli *ZSClient) ReconnectBaremetalPxeServer(uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.ReconnectBaremetalPxeServerEventView, error) {
	resp := view.ReconnectBaremetalPxeServerEventView{}
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

