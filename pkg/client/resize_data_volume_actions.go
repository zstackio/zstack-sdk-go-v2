// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResizeDataVolume operates on DataVolume
func (cli *ZSClient) ResizeDataVolume(uuid string, params param.ResizeDataVolumeParam) (*view.ResizeDataVolumeEventView, error) {
	resp := view.ResizeDataVolumeEventView{}
	if err := cli.Put("v1/volumes/data/resize/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
