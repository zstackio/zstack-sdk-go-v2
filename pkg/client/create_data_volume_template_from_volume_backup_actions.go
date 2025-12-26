// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolumeTemplateFromVolumeBackup creates DataVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeBackup(params param.CreateDataVolumeTemplateFromVolumeBackupParam) (*view.CreateDataVolumeTemplateFromVolumeBackupEventView, error) {
	resp := view.CreateDataVolumeTemplateFromVolumeBackupEventView{}
	if err := cli.Post("v1/images/data-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
