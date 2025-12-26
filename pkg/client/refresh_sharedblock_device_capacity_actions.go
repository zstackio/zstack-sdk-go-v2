// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshSharedblockDeviceCapacity operates on RefreshSharedblockDeviceCapacity
func (cli *ZSClient) RefreshSharedblockDeviceCapacity(params param.RefreshSharedblockDeviceCapacityParam) (*view.RefreshSharedBlockDeviceCapacityEventView, error) {
	resp := view.RefreshSharedBlockDeviceCapacityEventView{}
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
