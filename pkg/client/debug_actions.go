// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetDebugSignal gets DebugSignal by uuid
func (cli *ZSClient) GetDebugSignal() (*view.GetDebugSignalView, error) {
	var resp view.GetDebugSignalView
	if err := cli.GetWithRespKey("v1/debug", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DebugSignal operates on DebugSignal
func (cli *ZSClient) DebugSignal(params param.DebugSignalParam) (*view.DebugSignalEventView, error) {
	resp := view.DebugSignalEventView{}
	if err := cli.Post("v1/debug", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
