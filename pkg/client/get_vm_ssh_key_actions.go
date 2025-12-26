// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmSshKey gets VmSshKey by uuid
func (cli *ZSClient) GetVmSshKey(uuid string) (*view.GetVmSshKeyView, error) {
	var resp view.GetVmSshKeyView
	if err := cli.Get("v1/vm-instances/{uuid}/ssh-keys", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
