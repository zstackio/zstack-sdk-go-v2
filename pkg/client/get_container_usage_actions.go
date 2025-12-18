// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetContainerUsage gets ContainerUsage by uuid
func (cli *ZSClient) GetContainerUsage(uuid string) (*view.GetContainerUsageView, error) {
	var resp view.GetContainerUsageView
	if err := cli.Get("v1/container/usage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
