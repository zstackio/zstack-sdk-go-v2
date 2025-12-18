// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFaultToleranceVm 查询FaultToleranceVm列表
func (cli *ZSClient) QueryFaultToleranceVm(params param.QueryParam) ([]view.QueryFaultToleranceVmView, error) {
	var resp []view.QueryFaultToleranceVmView
	return resp, cli.List("v1/vm-instances/fault-tolerance", &params, &resp)
}

