// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBlockVolume updates BlockVolume
func (cli *ZSClient) UpdateBlockVolume(uuid string, params param.UpdateBlockVolumeParam) (*view.UpdateBlockVolumeEventView, error) {
	resp := view.UpdateBlockVolumeEventView{}
	if err := cli.Put("v1/block-volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
