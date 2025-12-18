// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncEcsVpcFromRemote 操作SyncEcsVpcFromRemote
func (cli *ZSClient) SyncEcsVpcFromRemote(params param.SyncEcsVpcFromRemoteParam) (*view.SyncEcsVpcFromRemoteEventView, error) {
	resp := view.SyncEcsVpcFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

