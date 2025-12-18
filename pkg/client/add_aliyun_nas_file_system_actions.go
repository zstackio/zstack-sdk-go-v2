// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunNasFileSystem 操作AddAliyunNasFileSystem
func (cli *ZSClient) AddAliyunNasFileSystem(uuid string, params param.AddAliyunNasFileSystemParam) (*view.AddAliyunNasFileSystemEventView, error) {
	resp := view.AddAliyunNasFileSystemEventView{}
	if err := cli.Put("v1/nas/aliyun", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

