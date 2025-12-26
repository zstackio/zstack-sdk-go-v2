// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmBootVolume operates on SetVmBootVolume
func (cli *ZSClient) SetVmBootVolume(uuid string, params param.SetVmBootVolumeParam) (*view.SetVmBootVolumeEventView, error) {
	resp := view.SetVmBootVolumeEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
