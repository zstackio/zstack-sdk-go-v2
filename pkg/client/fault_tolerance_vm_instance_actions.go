// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFaultToleranceVmInstance 创建FaultToleranceVmInstance
func (cli *ZSClient) CreateFaultToleranceVmInstance(params param.CreateFaultToleranceVmInstanceParam) (*view.CreateFaultToleranceVmInstanceEventView, error) {
	resp := view.CreateFaultToleranceVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/fault-tolerance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

