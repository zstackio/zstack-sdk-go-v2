// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunDisk updates AliyunDisk
func (cli *ZSClient) UpdateAliyunDisk(uuid string, params param.UpdateAliyunDiskParam) (*view.AliyunDiskInventoryView, error) {
	var resp view.UpdateAliyunDiskEventView
	if err := cli.Put("v1/hybrid/aliyun/disk/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
