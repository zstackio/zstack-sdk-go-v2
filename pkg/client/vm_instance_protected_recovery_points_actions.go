// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmInstanceProtectedRecoveryPoints 获取VmInstanceProtectedRecoveryPoints详情
func (cli *ZSClient) GetVmInstanceProtectedRecoveryPoints(uuid string) (*view.GetVmInstanceProtectedRecoveryPointsView, error) {
	var resp view.GetVmInstanceProtectedRecoveryPointsView
	if err := cli.Get("v1/vm-instances/{uuid}/protected-recovery-points", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

