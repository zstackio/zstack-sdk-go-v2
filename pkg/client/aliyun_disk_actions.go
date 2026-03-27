// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunDisk updates AliyunDisk
func (cli *ZSClient) UpdateAliyunDisk(ctx context.Context, uuid string, params param.UpdateAliyunDiskParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/disk", uuid, "", map[string]interface{}{
		"updateAliyunDisk": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
