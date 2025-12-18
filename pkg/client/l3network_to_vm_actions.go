// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachL3NetworkToVm 操作L3NetworkToVm
func (cli *ZSClient) AttachL3NetworkToVm(params param.AttachL3NetworkToVmParam) (*view.AttachL3NetworkToVmEventView, error) {
	resp := view.AttachL3NetworkToVmEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

