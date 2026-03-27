// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunSnapshot updates AliyunSnapshot
func (cli *ZSClient) UpdateAliyunSnapshot(ctx context.Context, uuid string, params param.UpdateAliyunSnapshotParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/snapshot", uuid, "", map[string]interface{}{
		"updateAliyunSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
