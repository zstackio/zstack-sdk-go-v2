// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddConnectionAccessPointFromRemote 操作AddConnectionAccessPointFromRemote
func (cli *ZSClient) AddConnectionAccessPointFromRemote(params param.AddConnectionAccessPointFromRemoteParam) (*view.AddConnectionAccessPointFromRemoteEventView, error) {
	resp := view.AddConnectionAccessPointFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/access-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

