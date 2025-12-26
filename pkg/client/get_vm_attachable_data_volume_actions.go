// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmAttachableDataVolume gets VmAttachableDataVolume by uuid
func (cli *ZSClient) GetVmAttachableDataVolume(uuid string) (*view.GetVmAttachableDataVolumeView, error) {
	var resp view.GetVmAttachableDataVolumeView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/data-volume-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
