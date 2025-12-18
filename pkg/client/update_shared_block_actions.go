// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSharedBlock updates SharedBlock
func (cli *ZSClient) UpdateSharedBlock(uuid string, params param.UpdateSharedBlockParam) (*view.UpdateSharedBlockEventView, error) {
	resp := view.UpdateSharedBlockEventView{}
	if err := cli.Put("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
