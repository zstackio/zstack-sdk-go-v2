// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetDebugSignal gets DebugSignal by uuid
func (cli *ZSClient) GetDebugSignal(uuid string) (*view.GetDebugSignalView, error) {
	var resp view.GetDebugSignalView
	if err := cli.Get("v1/debug", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
