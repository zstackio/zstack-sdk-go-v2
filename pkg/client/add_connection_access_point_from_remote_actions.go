// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddConnectionAccessPointFromRemote adds ConnectionAccessPointFromRemote
func (cli *ZSClient) AddConnectionAccessPointFromRemote(params param.AddConnectionAccessPointFromRemoteParam) (*view.AddConnectionAccessPointFromRemoteEventView, error) {
	resp := view.AddConnectionAccessPointFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/access-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
