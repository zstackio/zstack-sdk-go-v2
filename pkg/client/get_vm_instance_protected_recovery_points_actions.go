// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmInstanceProtectedRecoveryPoints gets VmInstanceProtectedRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceProtectedRecoveryPoints(uuid string) (*view.GetVmInstanceProtectedRecoveryPointsView, error) {
	var resp view.GetVmInstanceProtectedRecoveryPointsView
	if err := cli.Get("v1/vm-instances/{uuid}/protected-recovery-points", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
