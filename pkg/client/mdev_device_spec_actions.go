// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateMdevDeviceSpec 更新MdevDeviceSpec
func (cli *ZSClient) UpdateMdevDeviceSpec(uuid string, params param.UpdateMdevDeviceSpecParam) (*view.UpdateMdevDeviceSpecEventView, error) {
	resp := view.UpdateMdevDeviceSpecEventView{}
	if err := cli.Put("v1/mdev-device-specs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

