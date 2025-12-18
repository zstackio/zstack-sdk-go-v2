// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmRDP 操作SetVmRDP
func (cli *ZSClient) SetVmRDP(uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

