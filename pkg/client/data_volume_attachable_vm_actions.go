// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetDataVolumeAttachableVm 获取DataVolumeAttachableVm详情
func (cli *ZSClient) GetDataVolumeAttachableVm(uuid string) (*view.GetDataVolumeAttachableVmView, error) {
	var resp view.GetDataVolumeAttachableVmView
	if err := cli.Get("v1/volumes/{volumeUuid}/candidate-vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

