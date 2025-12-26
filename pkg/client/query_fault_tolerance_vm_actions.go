// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFaultToleranceVm queries FaultToleranceVm list
func (cli *ZSClient) QueryFaultToleranceVm(params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, error) {
	var resp []view.FaultToleranceVmGroupInventoryView
	return resp, cli.List("v1/vm-instances/fault-tolerance", params, &resp)
}
