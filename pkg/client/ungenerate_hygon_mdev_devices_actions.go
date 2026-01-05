// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UngenerateHygonMdevDevices operates on UngenerateHygonMdevDevices
func (cli *ZSClient) UngenerateHygonMdevDevices(uuid string, params param.UngenerateHygonMdevDevicesParam) (*view.UngenerateHygonMdevDevicesEventView, error) {
	resp := view.UngenerateHygonMdevDevicesEventView{}
	if err := cli.Put("v1/hygon-devices/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
