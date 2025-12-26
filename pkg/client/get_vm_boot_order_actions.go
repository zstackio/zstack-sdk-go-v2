// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmBootOrder gets VmBootOrder by uuid
func (cli *ZSClient) GetVmBootOrder(uuid string) (*view.GetVmBootOrderView, error) {
	var resp view.GetVmBootOrderView
	if err := cli.Get("v1/vm-instances/{uuid}/boot-orders", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
