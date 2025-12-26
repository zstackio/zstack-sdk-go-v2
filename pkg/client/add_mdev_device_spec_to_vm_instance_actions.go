// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddMdevDeviceSpecToVmInstance adds MdevDeviceSpecToVmInstance
func (cli *ZSClient) AddMdevDeviceSpecToVmInstance(params param.AddMdevDeviceSpecToVmInstanceParam) (*view.AddMdevDeviceSpecToVmInstanceEventView, error) {
	resp := view.AddMdevDeviceSpecToVmInstanceEventView{}
	if err := cli.Post("v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
