// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmNuma operates on SetVmNuma
func (cli *ZSClient) SetVmNuma(uuid string, params param.SetVmNumaParam) (*view.SetVmNumaEventView, error) {
	resp := view.SetVmNumaEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
