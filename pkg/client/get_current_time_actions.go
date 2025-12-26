// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCurrentTime gets CurrentTime by uuid
func (cli *ZSClient) GetCurrentTime(uuid string) (*view.GetCurrentTimeView, error) {
	var resp view.GetCurrentTimeView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
