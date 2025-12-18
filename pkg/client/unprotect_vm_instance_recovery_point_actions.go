// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnprotectVmInstanceRecoveryPoint 操作UnprotectVmInstanceRecoveryPoint
func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(uuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/unprotect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

