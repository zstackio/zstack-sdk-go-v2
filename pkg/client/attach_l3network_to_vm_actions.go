// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachL3NetworkToVm operates on L3NetworkToVm
func (cli *ZSClient) AttachL3NetworkToVm(params param.AttachL3NetworkToVmParam) (*view.AttachL3NetworkToVmEventView, error) {
	resp := view.AttachL3NetworkToVmEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
