// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVolumeBackup creates VolumeBackup
func (cli *ZSClient) CreateVolumeBackup(params param.CreateVolumeBackupParam) (*view.CreateVolumeBackupEventView, error) {
	resp := view.CreateVolumeBackupEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
