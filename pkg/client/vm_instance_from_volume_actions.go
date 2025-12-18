// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmInstanceFromVolume 创建VmInstanceFromVolume
func (cli *ZSClient) CreateVmInstanceFromVolume(params param.CreateVmInstanceFromVolumeParam) (*view.CreateVmInstanceFromVolumeEventView, error) {
	resp := view.CreateVmInstanceFromVolumeEventView{}
	if err := cli.Post("v1/vm-instances/from/volume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

