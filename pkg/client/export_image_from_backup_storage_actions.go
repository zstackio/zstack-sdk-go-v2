// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExportImageFromBackupStorage 操作ExportImageFromBackupStorage
func (cli *ZSClient) ExportImageFromBackupStorage(uuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

