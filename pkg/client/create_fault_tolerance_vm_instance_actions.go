// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFaultToleranceVmInstance creates FaultToleranceVmInstance
func (cli *ZSClient) CreateFaultToleranceVmInstance(params param.CreateFaultToleranceVmInstanceParam) (*view.CreateFaultToleranceVmInstanceEventView, error) {
	resp := view.CreateFaultToleranceVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/fault-tolerance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
