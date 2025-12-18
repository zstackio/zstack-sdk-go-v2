// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostResourceAllocation gets HostResourceAllocation by uuid
func (cli *ZSClient) GetHostResourceAllocation(uuid string) (*view.GetHostResourceAllocationEventView, error) {
	var resp view.GetHostResourceAllocationEventView
	if err := cli.Get("v1/hosts/{uuid}/resource-allocation", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
