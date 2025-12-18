// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateRootVolumeTemplateFromVolumeBackup creates RootVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeBackup(params param.CreateRootVolumeTemplateFromVolumeBackupParam) (*view.CreateRootVolumeTemplateFromVolumeBackupEventView, error) {
	resp := view.CreateRootVolumeTemplateFromVolumeBackupEventView{}
	if err := cli.Post("v1/images/root-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
