// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PutMetricData 操作PutMetricData
func (cli *ZSClient) PutMetricData(params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.Post("v1/zwatch/metrics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

