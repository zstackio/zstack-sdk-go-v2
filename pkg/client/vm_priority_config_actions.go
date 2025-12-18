// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmPriorityConfig 查询VmPriorityConfig列表
func (cli *ZSClient) QueryVmPriorityConfig(params param.QueryParam) ([]view.QueryVmPriorityConfigView, error) {
	var resp []view.QueryVmPriorityConfigView
	return resp, cli.List("v1/vm-priority-config", &params, &resp)
}

