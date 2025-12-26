// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetFactoryModeState gets FactoryModeState by uuid
func (cli *ZSClient) GetFactoryModeState(uuid string) (*view.GetFactoryModeStateView, error) {
	var resp view.GetFactoryModeStateView
	if err := cli.Get("v1/management-nodes/factory-mode-state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
