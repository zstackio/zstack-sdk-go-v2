// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSupportedCloudFormationResources gets SupportedCloudFormationResources by uuid
func (cli *ZSClient) GetSupportedCloudFormationResources(uuid string) (*view.GetSupportedCloudFormationResourcesView, error) {
	var resp view.GetSupportedCloudFormationResourcesView
	if err := cli.Get("v1/cloudformation/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
