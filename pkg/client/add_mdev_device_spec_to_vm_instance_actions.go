// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddMdevDeviceSpecToVmInstance adds MdevDeviceSpecToVmInstance
func (cli *ZSClient) AddMdevDeviceSpecToVmInstance(params param.AddMdevDeviceSpecToVmInstanceParam) (*view.AddMdevDeviceSpecToVmInstanceEventView, error) {
	resp := view.AddMdevDeviceSpecToVmInstanceEventView{}
	if err := cli.Post("v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
