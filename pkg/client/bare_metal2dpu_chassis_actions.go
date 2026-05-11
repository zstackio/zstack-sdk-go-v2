// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddBareMetal2DpuChassis adds BareMetal2DpuChassis
func (cli *ZSClient) AddBareMetal2DpuChassis(ctx context.Context, params param.AddBareMetal2DpuChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal2/chassis/dpu", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
