// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmFromCdpBackup 创建VmFromCdpBackup
func (cli *ZSClient) CreateVmFromCdpBackup(params param.CreateVmFromCdpBackupParam) (*view.CreateVmFromCdpBackupEventView, error) {
	resp := view.CreateVmFromCdpBackupEventView{}
	if err := cli.Post("v1/cdp-backups/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

