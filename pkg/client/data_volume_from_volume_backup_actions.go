// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataVolumeFromVolumeBackup 创建DataVolumeFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeFromVolumeBackup(params param.CreateDataVolumeFromVolumeBackupParam) (*view.CreateDataVolumeFromVolumeBackupEventView, error) {
	resp := view.CreateDataVolumeFromVolumeBackupEventView{}
	if err := cli.Post("v1/volumes/data-volume/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

