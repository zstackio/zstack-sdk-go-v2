// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GenerateHygonMdevDevices operates on GenerateHygonMdevDevices
func (cli *ZSClient) GenerateHygonMdevDevices(uuid string, params param.GenerateHygonMdevDevicesParam) (*view.GenerateHygonMdevDevicesEventView, error) {
	resp := view.GenerateHygonMdevDevicesEventView{}
	if err := cli.Put("v1/hygon-devices/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
