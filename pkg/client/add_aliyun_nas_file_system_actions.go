// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAliyunNasFileSystem adds AliyunNasFileSystem
func (cli *ZSClient) AddAliyunNasFileSystem(params param.AddAliyunNasFileSystemParam) (*view.AddAliyunNasFileSystemEventView, error) {
	resp := view.AddAliyunNasFileSystemEventView{}
	if err := cli.Post("v1/nas/aliyun", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
