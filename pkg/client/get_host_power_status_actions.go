// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostPowerStatus gets HostPowerStatus by uuid
func (cli *ZSClient) GetHostPowerStatus(uuid string) (*view.GetHostPowerStatusEventView, error) {
	var resp view.GetHostPowerStatusEventView
	if err := cli.Get("v1/hosts/power/{uuid}/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
