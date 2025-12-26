// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetDataVolumeAttachableVm gets DataVolumeAttachableVm by uuid
func (cli *ZSClient) GetDataVolumeAttachableVm(uuid string) (*view.GetDataVolumeAttachableVmView, error) {
	var resp view.GetDataVolumeAttachableVmView
	if err := cli.Get("v1/volumes/{volumeUuid}/candidate-vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
