// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPortMirrorSession queries PortMirrorSession list
func (cli *ZSClient) QueryPortMirrorSession(params *param.QueryParam) ([]view.PortMirrorSessionInventoryView, error) {
	var resp []view.PortMirrorSessionInventoryView
	return resp, cli.List("v1/port-mirrors/sessions", params, &resp)
}
