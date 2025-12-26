// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolumeFromVolumeBackup creates DataVolumeFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeFromVolumeBackup(params param.CreateDataVolumeFromVolumeBackupParam) (*view.CreateDataVolumeFromVolumeBackupEventView, error) {
	resp := view.CreateDataVolumeFromVolumeBackupEventView{}
	if err := cli.Post("v1/volumes/data-volume/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
