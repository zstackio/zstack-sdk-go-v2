// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2BondingNicRef 查询BareMetal2BondingNicRef列表
func (cli *ZSClient) QueryBareMetal2BondingNicRef(params param.QueryParam) ([]view.QueryBareMetal2ChassisView, error) {
	var resp []view.QueryBareMetal2ChassisView
	return resp, cli.List("v1/baremetal2/bonding/nic/refs", &params, &resp)
}

