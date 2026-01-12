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
	var resp view.UpdateAliyunSnapshotEventView
	err := cli.PutWithSpec("v1/hybrid/aliyun/snapshot", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
