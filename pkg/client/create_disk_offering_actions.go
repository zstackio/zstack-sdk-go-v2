// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDiskOffering creates DiskOffering
func (cli *ZSClient) CreateDiskOffering(params param.CreateDiskOfferingParam) (*view.CreateDiskOfferingEventView, error) {
	resp := view.CreateDiskOfferingEventView{}
	if err := cli.Post("v1/disk-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
