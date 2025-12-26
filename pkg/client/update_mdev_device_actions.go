// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateMdevDevice updates MdevDevice
func (cli *ZSClient) UpdateMdevDevice(uuid string, params param.UpdateMdevDeviceParam) (*view.UpdateMdevDeviceEventView, error) {
	resp := view.UpdateMdevDeviceEventView{}
	if err := cli.Put("v1/mdev-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
