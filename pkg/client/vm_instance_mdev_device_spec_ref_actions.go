// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmInstanceMdevDeviceSpecRef 查询VmInstanceMdevDeviceSpecRef列表
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(params param.QueryParam) ([]view.QueryVmInstanceMdevDeviceSpecRefView, error) {
	var resp []view.QueryVmInstanceMdevDeviceSpecRefView
	return resp, cli.List("v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", &params, &resp)
}

