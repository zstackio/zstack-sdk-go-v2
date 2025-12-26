// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBlockPrimaryStorage updates BlockPrimaryStorage
func (cli *ZSClient) UpdateBlockPrimaryStorage(uuid string, params param.UpdateBlockPrimaryStorageParam) (*view.UpdateBlockPrimaryStorageEventView, error) {
	resp := view.UpdateBlockPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/block/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
