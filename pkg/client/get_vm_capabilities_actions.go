// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmCapabilities gets VmCapabilities by uuid
func (cli *ZSClient) GetVmCapabilities(uuid string) (*view.GetVmCapabilitiesView, error) {
	var resp view.GetVmCapabilitiesView
	if err := cli.Get("v1/vm-instances/{uuid}/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
