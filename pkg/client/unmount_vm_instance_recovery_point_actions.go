// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UnmountVmInstanceRecoveryPoint operates on UnmountVmInstanceRecoveryPoint
func (cli *ZSClient) UnmountVmInstanceRecoveryPoint(params param.UnmountVmInstanceRecoveryPointParam) (*view.UnmountVmInstanceRecoveryPointEventView, error) {
	resp := view.UnmountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/unmount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
