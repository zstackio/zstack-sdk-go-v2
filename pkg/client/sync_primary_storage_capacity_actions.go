// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncPrimaryStorageCapacity 操作SyncPrimaryStorageCapacity
func (cli *ZSClient) SyncPrimaryStorageCapacity(uuid string, params param.SyncPrimaryStorageCapacityParam) (*view.SyncPrimaryStorageCapacityEventView, error) {
	resp := view.SyncPrimaryStorageCapacityEventView{}
	if err := cli.Put("v1/primary-storage/{primaryStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

