// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunNasFileSystem creates AliyunNasFileSystem
func (cli *ZSClient) CreateAliyunNasFileSystem(params param.CreateAliyunNasFileSystemParam) (*view.CreateNasFileSystemEventView, error) {
	resp := view.CreateNasFileSystemEventView{}
	if err := cli.Post("v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
