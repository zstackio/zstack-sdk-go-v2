// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmQga operates on SetVmQga
func (cli *ZSClient) SetVmQga(uuid string, params param.SetVmQgaParam) (*view.SetVmQgaEventView, error) {
	resp := view.SetVmQgaEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
