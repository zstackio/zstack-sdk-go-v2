// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetZWatchAlertHistogram 获取ZWatchAlertHistogram详情
func (cli *ZSClient) GetZWatchAlertHistogram(uuid string) (*view.GetZWatchAlertHistogramView, error) {
	var resp view.GetZWatchAlertHistogramView
	if err := cli.Get("v1/zwatch/alert-histories/histogram", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

