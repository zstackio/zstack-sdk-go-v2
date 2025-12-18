// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSharedBlockToSharedBlockGroup adds SharedBlockToSharedBlockGroup
func (cli *ZSClient) AddSharedBlockToSharedBlockGroup(params param.AddSharedBlockToSharedBlockGroupParam) (*view.AddSharedBlockToSharedBlockGroupEventView, error) {
	resp := view.AddSharedBlockToSharedBlockGroupEventView{}
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{uuid}/sharedblocks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
