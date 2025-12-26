// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateXskyBlockVolume updates XskyBlockVolume
func (cli *ZSClient) UpdateXskyBlockVolume(uuid string, params param.UpdateXskyBlockVolumeParam) (*view.UpdateBlockVolumeEventView, error) {
	resp := view.UpdateBlockVolumeEventView{}
	if err := cli.Put("v1/xsky/block-volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
