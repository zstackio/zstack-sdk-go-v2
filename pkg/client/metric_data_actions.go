// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMetricData 获取MetricData详情
func (cli *ZSClient) GetMetricData(uuid string) (*view.GetMetricDataView, error) {
	var resp view.GetMetricDataView
	if err := cli.Get("v1/zwatch/metrics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

