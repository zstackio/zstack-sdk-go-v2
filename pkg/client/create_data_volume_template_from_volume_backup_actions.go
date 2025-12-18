// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataVolumeTemplateFromVolumeBackup creates DataVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeBackup(params param.CreateDataVolumeTemplateFromVolumeBackupParam) (*view.CreateDataVolumeTemplateFromVolumeBackupEventView, error) {
	resp := view.CreateDataVolumeTemplateFromVolumeBackupEventView{}
	if err := cli.Post("v1/images/data-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
