// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSupportedCloudFormationResources 获取SupportedCloudFormationResources详情
func (cli *ZSClient) GetSupportedCloudFormationResources(uuid string) (*view.GetSupportedCloudFormationResourcesView, error) {
	var resp view.GetSupportedCloudFormationResourcesView
	if err := cli.Get("v1/cloudformation/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

