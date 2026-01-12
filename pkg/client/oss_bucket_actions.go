// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateOssBucket updates OssBucket
func (cli *ZSClient) UpdateOssBucket(uuid string, params param.UpdateOssBucketParam) (*view.OssBucketInventoryView, error) {
	var resp view.UpdateOssBucketEventView
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
