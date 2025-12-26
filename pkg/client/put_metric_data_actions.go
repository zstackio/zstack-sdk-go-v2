// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PutMetricData operates on PutMetricData
func (cli *ZSClient) PutMetricData(params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.Post("v1/zwatch/metrics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
