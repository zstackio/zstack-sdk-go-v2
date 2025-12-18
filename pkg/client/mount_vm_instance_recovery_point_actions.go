// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// MountVmInstanceRecoveryPoint 操作MountVmInstanceRecoveryPoint
func (cli *ZSClient) MountVmInstanceRecoveryPoint(params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/mount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

