// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UngenerateSeMdevDevices operates on UngenerateSeMdevDevices
func (cli *ZSClient) UngenerateSeMdevDevices(uuid string, params param.UngenerateSeMdevDevicesParam) (*view.UngenerateSeMdevDevicesEventView, error) {
	resp := view.UngenerateSeMdevDevicesEventView{}
	if err := cli.Put("v1/mtty-devices/{mttyDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
