// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmInstanceFromVolume creates VmInstanceFromVolume
func (cli *ZSClient) CreateVmInstanceFromVolume(params param.CreateVmInstanceFromVolumeParam) (*view.CreateVmInstanceFromVolumeEventView, error) {
	resp := view.CreateVmInstanceFromVolumeEventView{}
	if err := cli.Post("v1/vm-instances/from/volume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
