// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPortMirrorNetworkUsedIp queries PortMirrorNetworkUsedIp list
func (cli *ZSClient) QueryPortMirrorNetworkUsedIp(params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, error) {
	var resp []view.MirrorNetworkUsedIpInventoryView
	return resp, cli.List("v1/port-mirrors/networks/usedIps", params, &resp)
}
