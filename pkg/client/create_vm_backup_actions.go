// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmBackup creates VmBackup
func (cli *ZSClient) CreateVmBackup(params param.CreateVmBackupParam) (*view.CreateVmBackupEventView, error) {
	resp := view.CreateVmBackupEventView{}
	if err := cli.Post("v1/volumes/{rootVolumeUuid}/vm-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
