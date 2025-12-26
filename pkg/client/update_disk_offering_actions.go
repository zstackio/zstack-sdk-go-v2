// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateDiskOffering updates DiskOffering
func (cli *ZSClient) UpdateDiskOffering(uuid string, params param.UpdateDiskOfferingParam) (*view.UpdateDiskOfferingEventView, error) {
	resp := view.UpdateDiskOfferingEventView{}
	if err := cli.Put("v1/disk-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
