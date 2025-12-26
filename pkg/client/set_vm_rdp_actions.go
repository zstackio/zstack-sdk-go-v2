// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmRDP operates on SetVmRDP
func (cli *ZSClient) SetVmRDP(uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
