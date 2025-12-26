// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// MountVmInstanceRecoveryPoint operates on MountVmInstanceRecoveryPoint
func (cli *ZSClient) MountVmInstanceRecoveryPoint(params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/mount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
