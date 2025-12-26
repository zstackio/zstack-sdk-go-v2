// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmInstanceHaLevel gets VmInstanceHaLevel by uuid
func (cli *ZSClient) GetVmInstanceHaLevel(uuid string) (*view.GetVmInstanceHaLevelView, error) {
	var resp view.GetVmInstanceHaLevelView
	if err := cli.Get("v1/vm-instances/{uuid}/ha-levels", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
