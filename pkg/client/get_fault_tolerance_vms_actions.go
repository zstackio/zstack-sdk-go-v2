// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetFaultToleranceVms gets FaultToleranceVms by uuid
func (cli *ZSClient) GetFaultToleranceVms(uuid string) (*view.GetFaultToleranceVmsView, error) {
	var resp view.GetFaultToleranceVmsView
	if err := cli.Get("v1/vm-instances/fault-tolerance/sub-vms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
