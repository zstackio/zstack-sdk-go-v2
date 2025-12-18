// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateMdevDevice 更新MdevDevice
func (cli *ZSClient) UpdateMdevDevice(uuid string, params param.UpdateMdevDeviceParam) (*view.UpdateMdevDeviceEventView, error) {
	resp := view.UpdateMdevDeviceEventView{}
	if err := cli.Put("v1/mdev-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

