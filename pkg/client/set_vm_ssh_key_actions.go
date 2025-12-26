// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmSshKey operates on SetVmSshKey
func (cli *ZSClient) SetVmSshKey(uuid string, params param.SetVmSshKeyParam) (*view.SetVmSshKeyEventView, error) {
	resp := view.SetVmSshKeyEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
