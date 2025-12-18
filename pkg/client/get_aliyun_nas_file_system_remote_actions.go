// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAliyunNasFileSystemRemote gets AliyunNasFileSystemRemote by uuid
func (cli *ZSClient) GetAliyunNasFileSystemRemote(uuid string) (*view.GetAliyunNasFileSystemRemoteView, error) {
	var resp view.GetAliyunNasFileSystemRemoteView
	if err := cli.Get("v1/nas/aliyun/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
