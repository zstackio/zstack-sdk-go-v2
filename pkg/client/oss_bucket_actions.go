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
	err := cli.PutWithSpec("v1/hybrid/aliyun/oss-bucket", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
