// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetConnectionAccessPointFromRemote gets ConnectionAccessPointFromRemote by uuid
func (cli *ZSClient) GetConnectionAccessPointFromRemote(uuid string) (*view.GetConnectionAccessPointFromRemoteView, error) {
	var resp view.GetConnectionAccessPointFromRemoteView
	if err := cli.Get("v1/hybrid/aliyun/access-point{dataCenterUuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
