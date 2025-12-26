// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachDataVolumeToVm operates on DataVolumeToVm
func (cli *ZSClient) AttachDataVolumeToVm(params param.AttachDataVolumeToVmParam) (*view.AttachDataVolumeToVmEventView, error) {
	resp := view.AttachDataVolumeToVmEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
