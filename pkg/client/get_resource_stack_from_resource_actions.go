// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetResourceStackFromResource gets ResourceStackFromResource by uuid
func (cli *ZSClient) GetResourceStackFromResource(uuid string) (*view.GetResourceStackFromResourceView, error) {
	var resp view.GetResourceStackFromResourceView
	if err := cli.Get("v1/cloudformation/resources/stack", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
