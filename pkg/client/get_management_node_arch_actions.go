// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetManagementNodeArch gets ManagementNodeArch by uuid
func (cli *ZSClient) GetManagementNodeArch(uuid string) (*view.GetManagementNodeArchView, error) {
	var resp view.GetManagementNodeArchView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
