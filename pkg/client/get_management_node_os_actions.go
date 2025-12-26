// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetManagementNodeOS gets ManagementNodeOS by uuid
func (cli *ZSClient) GetManagementNodeOS(uuid string) (*view.GetManagementNodeOSView, error) {
	var resp view.GetManagementNodeOSView
	if err := cli.Get("v1/management/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
