// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVersion gets Version by uuid
func (cli *ZSClient) GetVersion(uuid string) (*view.GetVersionView, error) {
	var resp view.GetVersionView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
