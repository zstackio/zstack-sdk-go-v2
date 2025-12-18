// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachVmNicToVm operates on VmNicToVm
func (cli *ZSClient) AttachVmNicToVm(params param.AttachVmNicToVmParam) (*view.AttachVmNicToVmEventView, error) {
	resp := view.AttachVmNicToVmEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/nices/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
