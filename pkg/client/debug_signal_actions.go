// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DebugSignal operates on DebugSignal
func (cli *ZSClient) DebugSignal(params param.DebugSignalParam) (*view.DebugSignalEventView, error) {
	resp := view.DebugSignalEventView{}
	if err := cli.Post("v1/debug", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
