// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BackupStorageMigrateImage operates on BackupStorageMigrateImage
func (cli *ZSClient) BackupStorageMigrateImage(uuid string, params param.BackupStorageMigrateImageParam) (*view.BackupStorageMigrateImageEventView, error) {
	resp := view.BackupStorageMigrateImageEventView{}
	if err := cli.Put("v1/backup-storage/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
