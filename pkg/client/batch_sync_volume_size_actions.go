// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BatchSyncVolumeSize 操作BatchSyncVolumeSize
func (cli *ZSClient) BatchSyncVolumeSize(params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.Post("v1/volumes/batch-sync-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

