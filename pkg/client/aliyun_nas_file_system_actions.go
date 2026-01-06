// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasFileSystem adds AliyunNasFileSystem
func (cli *ZSClient) AddAliyunNasFileSystem(params param.AddAliyunNasFileSystemParam) (*view.AliyunNasFileSystemInventoryView, error) {
	var resp view.AddAliyunNasFileSystemEventView
	if err := cli.Post("v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateAliyunNasFileSystem creates AliyunNasFileSystem
func (cli *ZSClient) CreateAliyunNasFileSystem(params param.CreateAliyunNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	var resp view.CreateNasFileSystemEventView
	if err := cli.Post("v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
