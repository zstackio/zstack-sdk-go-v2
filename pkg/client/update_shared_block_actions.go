// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSharedBlock updates SharedBlock
func (cli *ZSClient) UpdateSharedBlock(uuid string, params param.UpdateSharedBlockParam) (*view.UpdateSharedBlockEventView, error) {
	resp := view.UpdateSharedBlockEventView{}
	if err := cli.Put("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
