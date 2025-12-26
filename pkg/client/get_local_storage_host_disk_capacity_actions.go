// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLocalStorageHostDiskCapacity gets LocalStorageHostDiskCapacity by uuid
func (cli *ZSClient) GetLocalStorageHostDiskCapacity(uuid string) (*view.GetLocalStorageHostDiskCapacityView, error) {
	var resp view.GetLocalStorageHostDiskCapacityView
	if err := cli.Get("v1/primary-storage/local-storage/{primaryStorageUuid}/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
