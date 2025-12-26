// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSharedMountPointPrimaryStorage adds SharedMountPointPrimaryStorage
func (cli *ZSClient) AddSharedMountPointPrimaryStorage(params param.AddSharedMountPointPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/smp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
