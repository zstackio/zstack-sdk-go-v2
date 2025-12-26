// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVolume updates Volume
func (cli *ZSClient) UpdateVolume(uuid string, params param.UpdateVolumeParam) (*view.UpdateVolumeEventView, error) {
	resp := view.UpdateVolumeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
