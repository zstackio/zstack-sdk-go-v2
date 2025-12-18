// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPhysicalSwitch 查询PhysicalSwitch列表
func (cli *ZSClient) QueryPhysicalSwitch(params param.QueryParam) ([]view.QueryPhysicalSwitchView, error) {
	var resp []view.QueryPhysicalSwitchView
	return resp, cli.List("v1/topo/physical-switches", &params, &resp)
}

