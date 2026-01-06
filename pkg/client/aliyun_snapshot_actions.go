// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunSnapshot updates AliyunSnapshot
func (cli *ZSClient) UpdateAliyunSnapshot(uuid string, params param.UpdateAliyunSnapshotParam) (*view.AliyunSnapshotInventoryView, error) {
	var resp view.UpdateAliyunSnapshotEventView
	if err := cli.Put("v1/hybrid/aliyun/snapshot/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
