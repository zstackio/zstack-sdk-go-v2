// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetResourceFromResourceStack gets ResourceFromResourceStack by uuid
func (cli *ZSClient) GetResourceFromResourceStack(uuid string) (*view.GetResourceFromResourceStackView, error) {
	var resp view.GetResourceFromResourceStackView
	if err := cli.Get("v1/cloudformation/stack/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
