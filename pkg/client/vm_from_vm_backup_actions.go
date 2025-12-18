// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmFromVmBackup 创建VmFromVmBackup
func (cli *ZSClient) CreateVmFromVmBackup(params param.CreateVmFromVmBackupParam) (*view.CreateVmFromVmBackupEventView, error) {
	resp := view.CreateVmFromVmBackupEventView{}
	if err := cli.Post("v1/vm-instances/from/vm-backups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

