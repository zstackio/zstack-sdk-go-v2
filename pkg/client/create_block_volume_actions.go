// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBlockVolume creates BlockVolume
func (cli *ZSClient) CreateBlockVolume(params param.CreateBlockVolumeParam) (*view.CreateBlockVolumeEventView, error) {
	resp := view.CreateBlockVolumeEventView{}
	if err := cli.Post("v1/block-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
