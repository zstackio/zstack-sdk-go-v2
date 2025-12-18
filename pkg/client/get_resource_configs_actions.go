// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceConfigs gets ResourceConfigs by uuid
func (cli *ZSClient) GetResourceConfigs(uuid string) (*view.GetResourceConfigsView, error) {
	var resp view.GetResourceConfigsView
	if err := cli.Get("v1/resource-configurations/{resourceUuid}/{category}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
