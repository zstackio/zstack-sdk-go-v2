// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RebootVmInstance operates on VmInstance
func (cli *ZSClient) RebootVmInstance(uuid string, params param.RebootVmInstanceParam) (*view.RebootVmInstanceEventView, error) {
	resp := view.RebootVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
