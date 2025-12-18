// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetTrashOnBackupStorage gets TrashOnBackupStorage by uuid
func (cli *ZSClient) GetTrashOnBackupStorage(uuid string) (*view.GetTrashOnBackupStorageView, error) {
	var resp view.GetTrashOnBackupStorageView
	if err := cli.Get("v1/backup-storage/trash", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
