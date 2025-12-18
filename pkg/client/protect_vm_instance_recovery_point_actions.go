// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ProtectVmInstanceRecoveryPoint 操作ProtectVmInstanceRecoveryPoint
func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(uuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/protect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

