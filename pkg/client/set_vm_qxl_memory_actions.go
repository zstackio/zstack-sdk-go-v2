// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmQxlMemory operates on SetVmQxlMemory
func (cli *ZSClient) SetVmQxlMemory(uuid string, params param.SetVmQxlMemoryParam) (*view.SetVmQxlMemoryEventView, error) {
	resp := view.SetVmQxlMemoryEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
