// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StopBaremetalPxeServer 停止BaremetalPxeServer
func (cli *ZSClient) StopBaremetalPxeServer(uuid string, params param.StopBaremetalPxeServerParam) (*view.StopBaremetalPxeServerEventView, error) {
	resp := view.StopBaremetalPxeServerEventView{}
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

