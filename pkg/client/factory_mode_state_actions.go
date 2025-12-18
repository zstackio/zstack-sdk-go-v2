// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetFactoryModeState 获取FactoryModeState详情
func (cli *ZSClient) GetFactoryModeState(uuid string) (*view.GetFactoryModeStateView, error) {
	var resp view.GetFactoryModeStateView
	if err := cli.Get("v1/management-nodes/factory-mode-state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

