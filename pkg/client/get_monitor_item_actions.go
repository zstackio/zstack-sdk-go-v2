// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMonitorItem gets MonitorItem by uuid
func (cli *ZSClient) GetMonitorItem(uuid string) (*view.GetMonitorItemView, error) {
	var resp view.GetMonitorItemView
	if err := cli.Get("v1/monitoring/items", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
