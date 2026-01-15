// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasFileSystem adds AliyunNasFileSystem
func (cli *ZSClient) AddAliyunNasFileSystem(params param.AddAliyunNasFileSystemParam) (*view.AliyunNasFileSystemInventoryView, error) {
	resp := view.AliyunNasFileSystemInventoryView{}
	if err := cli.Post("v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAliyunNasFileSystem creates AliyunNasFileSystem
func (cli *ZSClient) CreateAliyunNasFileSystem(params param.CreateAliyunNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	resp := view.NasFileSystemInventoryView{}
	if err := cli.Post("v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
