// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBlockPrimaryStorageMetadata 获取BlockPrimaryStorageMetadata详情
func (cli *ZSClient) GetBlockPrimaryStorageMetadata(uuid string) (*view.QueryBlockPrimaryStorageView, error) {
	var resp view.QueryBlockPrimaryStorageView
	if err := cli.Get("v1/primary-storage/block/metadata", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

