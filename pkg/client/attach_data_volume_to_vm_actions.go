// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachDataVolumeToVm operates on DataVolumeToVm
func (cli *ZSClient) AttachDataVolumeToVm(params param.AttachDataVolumeToVmParam) (*view.AttachDataVolumeToVmEventView, error) {
	resp := view.AttachDataVolumeToVmEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
