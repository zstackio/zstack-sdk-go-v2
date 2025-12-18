// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEcsVSwitchFromLocal 查询EcsVSwitchFromLocal列表
func (cli *ZSClient) QueryEcsVSwitchFromLocal(params param.QueryParam) ([]view.QueryEcsVSwitchFromLocalView, error) {
	var resp []view.QueryEcsVSwitchFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/vswitch", &params, &resp)
}

