// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetMonitorItem gets MonitorItem by uuid
func (cli *ZSClient) GetMonitorItem(uuid string) (*view.GetMonitorItemView, error) {
	var resp view.GetMonitorItemView
	if err := cli.Get("v1/monitoring/items", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
