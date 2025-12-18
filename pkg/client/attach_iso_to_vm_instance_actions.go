// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachIsoToVmInstance operates on IsoToVmInstance
func (cli *ZSClient) AttachIsoToVmInstance(params param.AttachIsoToVmInstanceParam) (*view.AttachIsoToVmInstanceEventView, error) {
	resp := view.AttachIsoToVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/iso/{isoUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
