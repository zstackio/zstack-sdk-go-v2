// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPortMirrorNetworkUsedIp 查询PortMirrorNetworkUsedIp列表
func (cli *ZSClient) QueryPortMirrorNetworkUsedIp(params param.QueryParam) ([]view.QueryPortMirrorNetworkUsedIpView, error) {
	var resp []view.QueryPortMirrorNetworkUsedIpView
	return resp, cli.List("v1/port-mirrors/networks/usedIps", &params, &resp)
}

