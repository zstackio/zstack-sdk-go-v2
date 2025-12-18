// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetDebugSignal gets DebugSignal by uuid
func (cli *ZSClient) GetDebugSignal(uuid string) (*view.GetDebugSignalView, error) {
	var resp view.GetDebugSignalView
	if err := cli.Get("v1/debug", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
