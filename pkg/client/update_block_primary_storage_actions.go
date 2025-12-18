// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBlockPrimaryStorage updates BlockPrimaryStorage
func (cli *ZSClient) UpdateBlockPrimaryStorage(uuid string, params param.UpdateBlockPrimaryStorageParam) (*view.UpdateBlockPrimaryStorageEventView, error) {
	resp := view.UpdateBlockPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/block/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
