// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetFaultToleranceVms 获取FaultToleranceVms详情
func (cli *ZSClient) GetFaultToleranceVms(uuid string) (*view.GetFaultToleranceVmsView, error) {
	var resp view.GetFaultToleranceVmsView
	if err := cli.Get("v1/vm-instances/fault-tolerance/sub-vms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

