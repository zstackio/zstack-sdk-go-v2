// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBlockPrimaryStorageMetadata gets BlockPrimaryStorageMetadata by uuid
func (cli *ZSClient) GetBlockPrimaryStorageMetadata(uuid string) (*view.QueryBlockPrimaryStorageView, error) {
	var resp view.QueryBlockPrimaryStorageView
	if err := cli.Get("v1/primary-storage/block/metadata", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
