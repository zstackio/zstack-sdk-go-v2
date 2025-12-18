// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHostNetworkInterfaceLldp 查询HostNetworkInterfaceLldp列表
func (cli *ZSClient) QueryHostNetworkInterfaceLldp(params param.QueryParam) ([]view.QueryHostNetworkInterfaceLldpView, error) {
	var resp []view.QueryHostNetworkInterfaceLldpView
	return resp, cli.List("v1/hostNetworkInterface/lldp/all", &params, &resp)
}

