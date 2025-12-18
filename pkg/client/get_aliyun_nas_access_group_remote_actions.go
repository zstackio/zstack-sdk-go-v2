// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAliyunNasAccessGroupRemote gets AliyunNasAccessGroupRemote by uuid
func (cli *ZSClient) GetAliyunNasAccessGroupRemote(uuid string) (*view.GetAliyunNasAccessGroupRemoteView, error) {
	var resp view.GetAliyunNasAccessGroupRemoteView
	if err := cli.Get("v1/nas/aliyun/access/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
