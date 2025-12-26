// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmEmulatorPinning operates on SetVmEmulatorPinning
func (cli *ZSClient) SetVmEmulatorPinning(uuid string, params param.SetVmEmulatorPinningParam) (*view.SetVmEmulatorPinningEventView, error) {
	resp := view.SetVmEmulatorPinningEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
