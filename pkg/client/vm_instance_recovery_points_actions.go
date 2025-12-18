// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmInstanceRecoveryPoints 获取VmInstanceRecoveryPoints详情
func (cli *ZSClient) GetVmInstanceRecoveryPoints(uuid string) (*view.GetVmInstanceRecoveryPointsView, error) {
	var resp view.GetVmInstanceRecoveryPointsView
	if err := cli.Get("v1/vm-instances/{uuid}/recovery-points", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

