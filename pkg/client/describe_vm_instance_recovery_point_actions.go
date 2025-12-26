// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DescribeVmInstanceRecoveryPoint operates on DescribeVmInstanceRecoveryPoint
func (cli *ZSClient) DescribeVmInstanceRecoveryPoint(params param.DescribeVmInstanceRecoveryPointParam) (*view.DescribeVmInstanceRecoveryPointView, error) {
	var resp view.DescribeVmInstanceRecoveryPointView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/recovery-point", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
