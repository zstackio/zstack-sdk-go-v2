// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPortMirror queries PortMirror list
func (cli *ZSClient) QueryPortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, error) {
	var resp []view.PortMirrorInventoryView
	return resp, cli.List("v1/port-mirrors", params, &resp)
}
