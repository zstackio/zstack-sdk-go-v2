// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmInstanceRecoveryPoints gets VmInstanceRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceRecoveryPoints(uuid string) (*view.GetVmInstanceRecoveryPointsView, error) {
	var resp view.GetVmInstanceRecoveryPointsView
	if err := cli.Get("v1/vm-instances/{uuid}/recovery-points", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
