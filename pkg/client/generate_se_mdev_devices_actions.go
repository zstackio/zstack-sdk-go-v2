// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GenerateSeMdevDevices operates on GenerateSeMdevDevices
func (cli *ZSClient) GenerateSeMdevDevices(uuid string, params param.GenerateSeMdevDevicesParam) (*view.GenerateSeMdevDevicesEventView, error) {
	resp := view.GenerateSeMdevDevicesEventView{}
	if err := cli.Put("v1/mtty-devices/{mttyDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
