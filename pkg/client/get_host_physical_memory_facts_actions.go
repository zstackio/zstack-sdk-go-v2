// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostPhysicalMemoryFacts gets HostPhysicalMemoryFacts by uuid
func (cli *ZSClient) GetHostPhysicalMemoryFacts(uuid string) (*view.GetHostPhysicalMemoryFactsView, error) {
	var resp view.GetHostPhysicalMemoryFactsView
	if err := cli.Get("v1/hosts/physical-memory-facts/{hostUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
