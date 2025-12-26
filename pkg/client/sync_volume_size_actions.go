// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVolumeSize operates on SyncVolumeSize
func (cli *ZSClient) SyncVolumeSize(uuid string, params param.SyncVolumeSizeParam) (*view.SyncVolumeSizeEventView, error) {
	resp := view.SyncVolumeSizeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
