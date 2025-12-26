// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachVmNicToVm operates on VmNicToVm
func (cli *ZSClient) AttachVmNicToVm(params param.AttachVmNicToVmParam) (*view.AttachVmNicToVmEventView, error) {
	resp := view.AttachVmNicToVmEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/nices/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
