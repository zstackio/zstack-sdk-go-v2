// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmBootVolume operates on SetVmBootVolume
func (cli *ZSClient) SetVmBootVolume(uuid string, params param.SetVmBootVolumeParam) (*view.SetVmBootVolumeEventView, error) {
	resp := view.SetVmBootVolumeEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
