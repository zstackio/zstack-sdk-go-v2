// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetConnectionAccessPointFromRemote 获取ConnectionAccessPointFromRemote详情
func (cli *ZSClient) GetConnectionAccessPointFromRemote(uuid string) (*view.GetConnectionAccessPointFromRemoteView, error) {
	var resp view.GetConnectionAccessPointFromRemoteView
	if err := cli.Get("v1/hybrid/aliyun/access-point{dataCenterUuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

