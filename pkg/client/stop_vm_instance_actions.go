// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StopVmInstance stops VmInstance
func (cli *ZSClient) StopVmInstance(uuid string, params param.StopVmInstanceParam) (*view.StopVmInstanceEventView, error) {
	resp := view.StopVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
