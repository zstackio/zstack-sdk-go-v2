// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmsCapabilities gets VmsCapabilities by uuid
func (cli *ZSClient) GetVmsCapabilities(uuid string) (*view.GetVmsCapabilitiesView, error) {
	var resp view.GetVmsCapabilitiesView
	if err := cli.Get("v1/vm-instances/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
