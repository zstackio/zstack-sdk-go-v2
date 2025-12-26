// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ProtectVmInstanceRecoveryPoint operates on ProtectVmInstanceRecoveryPoint
func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(uuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/protect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
