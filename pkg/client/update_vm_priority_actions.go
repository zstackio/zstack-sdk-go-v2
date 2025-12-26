// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVmPriority updates VmPriority
func (cli *ZSClient) UpdateVmPriority(uuid string, params param.UpdateVmPriorityParam) (*view.UpdateVmPriorityEventView, error) {
	resp := view.UpdateVmPriorityEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
