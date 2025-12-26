// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BatchSyncVolumeSize operates on BatchSyncVolumeSize
func (cli *ZSClient) BatchSyncVolumeSize(params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.Post("v1/volumes/batch-sync-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
