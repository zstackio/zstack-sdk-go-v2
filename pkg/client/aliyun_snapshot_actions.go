// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunSnapshot updates AliyunSnapshot
func (cli *ZSClient) UpdateAliyunSnapshot(uuid string, params param.UpdateAliyunSnapshotParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/snapshot", uuid, map[string]interface{}{
		"updateAliyunSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
