// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMonitorItem 获取MonitorItem详情
func (cli *ZSClient) GetMonitorItem(uuid string) (*view.GetMonitorItemView, error) {
	var resp view.GetMonitorItemView
	if err := cli.Get("v1/monitoring/items", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

