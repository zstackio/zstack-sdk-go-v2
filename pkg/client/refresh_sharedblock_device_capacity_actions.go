// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshSharedblockDeviceCapacity 操作RefreshSharedblockDeviceCapacity
func (cli *ZSClient) RefreshSharedblockDeviceCapacity(params param.RefreshSharedblockDeviceCapacityParam) (*view.RefreshSharedBlockDeviceCapacityEventView, error) {
	resp := view.RefreshSharedBlockDeviceCapacityEventView{}
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

