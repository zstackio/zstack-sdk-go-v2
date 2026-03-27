// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateOssBucket updates OssBucket
func (cli *ZSClient) UpdateOssBucket(ctx context.Context, uuid string, params param.UpdateOssBucketParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/oss-bucket", uuid, "", map[string]interface{}{
		"updateOssBucket": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
