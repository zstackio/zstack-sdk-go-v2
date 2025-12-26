// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// FlattenVmInstance operates on FlattenVmInstance
func (cli *ZSClient) FlattenVmInstance(uuid string, params param.FlattenVmInstanceParam) (*view.FlattenVmInstanceEventView, error) {
	resp := view.FlattenVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
