// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateDiskOffering updates DiskOffering
func (cli *ZSClient) UpdateDiskOffering(uuid string, params param.UpdateDiskOfferingParam) (*view.UpdateDiskOfferingEventView, error) {
	resp := view.UpdateDiskOfferingEventView{}
	if err := cli.Put("v1/disk-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
