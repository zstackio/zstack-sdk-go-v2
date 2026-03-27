// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasFileSystem adds AliyunNasFileSystem
func (cli *ZSClient) AddAliyunNasFileSystem(ctx context.Context, params param.AddAliyunNasFileSystemParam) (*view.AliyunNasFileSystemInventoryView, error) {
	resp := view.AliyunNasFileSystemInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/nas/aliyun", "", "", map[string]interface{}{
		"addAliyunNasFileSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAliyunNasFileSystem creates AliyunNasFileSystem
func (cli *ZSClient) CreateAliyunNasFileSystem(ctx context.Context, params param.CreateAliyunNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	resp := view.NasFileSystemInventoryView{}
	if err := cli.Post(ctx, "v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
