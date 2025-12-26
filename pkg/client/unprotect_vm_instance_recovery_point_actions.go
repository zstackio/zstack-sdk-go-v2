// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UnprotectVmInstanceRecoveryPoint operates on UnprotectVmInstanceRecoveryPoint
func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(uuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/unprotect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
