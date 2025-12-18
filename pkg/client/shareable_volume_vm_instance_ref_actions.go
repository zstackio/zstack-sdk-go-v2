// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryShareableVolumeVmInstanceRef 查询ShareableVolumeVmInstanceRef列表
func (cli *ZSClient) QueryShareableVolumeVmInstanceRef(params param.QueryParam) ([]view.QueryShareableVolumeVmInstanceRefView, error) {
	var resp []view.QueryShareableVolumeVmInstanceRefView
	return resp, cli.List("v1/volumes/vm-instances/refs", &params, &resp)
}

