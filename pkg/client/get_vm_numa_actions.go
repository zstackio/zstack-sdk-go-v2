// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmNuma gets VmNuma by uuid
func (cli *ZSClient) GetVmNuma(uuid string) (*view.GetVmNumaView, error) {
	var resp view.GetVmNumaView
	if err := cli.Get("v1/vm-instances/{uuid}/vnuma", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
