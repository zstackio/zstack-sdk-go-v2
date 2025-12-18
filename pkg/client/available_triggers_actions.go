// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAvailableTriggers 获取AvailableTriggers详情
func (cli *ZSClient) GetAvailableTriggers(uuid string) (*view.GetAvailableTriggersView, error) {
	var resp view.GetAvailableTriggersView
	if err := cli.Get("v1/scheduler/triggers/available", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

