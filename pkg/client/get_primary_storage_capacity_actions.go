// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageCapacity gets PrimaryStorageCapacity by uuid
func (cli *ZSClient) GetPrimaryStorageCapacity(uuid string) (*view.GetPrimaryStorageCapacityView, error) {
	var resp view.GetPrimaryStorageCapacityView
	if err := cli.Get("v1/primary-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
