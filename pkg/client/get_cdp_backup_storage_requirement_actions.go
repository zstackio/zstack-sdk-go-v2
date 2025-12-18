// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCdpBackupStorageRequirement gets CdpBackupStorageRequirement by uuid
func (cli *ZSClient) GetCdpBackupStorageRequirement(uuid string) (*view.GetCdpBackupStorageRequirementView, error) {
	var resp view.GetCdpBackupStorageRequirementView
	if err := cli.Get("v1/cdp-backup-storage/{backupStorageUuid}/requirement", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
