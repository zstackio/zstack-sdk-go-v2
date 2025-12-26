// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResizeRootVolume operates on RootVolume
func (cli *ZSClient) ResizeRootVolume(uuid string, params param.ResizeRootVolumeParam) (*view.ResizeRootVolumeEventView, error) {
	resp := view.ResizeRootVolumeEventView{}
	if err := cli.Put("v1/volumes/resize/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
