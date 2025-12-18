// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmAttachableDataVolume 获取VmAttachableDataVolume详情
func (cli *ZSClient) GetVmAttachableDataVolume(uuid string) (*view.GetVmAttachableDataVolumeView, error) {
	var resp view.GetVmAttachableDataVolumeView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/data-volume-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

