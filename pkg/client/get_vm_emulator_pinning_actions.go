// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmEmulatorPinning gets VmEmulatorPinning by uuid
func (cli *ZSClient) GetVmEmulatorPinning(uuid string) (*view.GetVmEmulatorPinningView, error) {
	var resp view.GetVmEmulatorPinningView
	if err := cli.Get("v1/vm-instances/{uuid}/emulator-pinning", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
