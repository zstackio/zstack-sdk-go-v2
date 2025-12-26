// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverVmInstance operates on VmInstance
func (cli *ZSClient) RecoverVmInstance(uuid string, params param.RecoverVmInstanceParam) (*view.RecoverVmInstanceEventView, error) {
	resp := view.RecoverVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
