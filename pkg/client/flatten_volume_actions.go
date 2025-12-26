// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// FlattenVolume operates on FlattenVolume
func (cli *ZSClient) FlattenVolume(uuid string, params param.FlattenVolumeParam) (*view.FlattenVolumeEventView, error) {
	resp := view.FlattenVolumeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
