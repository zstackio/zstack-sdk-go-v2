// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmBootOrder operates on SetVmBootOrder
func (cli *ZSClient) SetVmBootOrder(uuid string, params param.SetVmBootOrderParam) (*view.SetVmBootOrderEventView, error) {
	resp := view.SetVmBootOrderEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
