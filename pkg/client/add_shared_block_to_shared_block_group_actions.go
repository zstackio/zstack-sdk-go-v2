// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSharedBlockToSharedBlockGroup adds SharedBlockToSharedBlockGroup
func (cli *ZSClient) AddSharedBlockToSharedBlockGroup(params param.AddSharedBlockToSharedBlockGroupParam) (*view.AddSharedBlockToSharedBlockGroupEventView, error) {
	resp := view.AddSharedBlockToSharedBlockGroupEventView{}
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{uuid}/sharedblocks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
