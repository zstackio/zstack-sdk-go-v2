// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSharedBlockGroupPrimaryStorage adds SharedBlockGroupPrimaryStorage
func (cli *ZSClient) AddSharedBlockGroupPrimaryStorage(params param.AddSharedBlockGroupPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/sharedblockgroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
