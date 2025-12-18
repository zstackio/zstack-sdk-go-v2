// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// FailoverFaultToleranceVm operates on FailoverFaultToleranceVm
func (cli *ZSClient) FailoverFaultToleranceVm(uuid string, params param.FailoverFaultToleranceVmParam) (*view.FailoverFaultToleranceVmEventView, error) {
	resp := view.FailoverFaultToleranceVmEventView{}
	if err := cli.Put("v1/vm-instances/fault-tolerance", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
