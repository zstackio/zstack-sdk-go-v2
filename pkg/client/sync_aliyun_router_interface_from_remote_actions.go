// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncAliyunRouterInterfaceFromRemote 操作SyncAliyunRouterInterfaceFromRemote
func (cli *ZSClient) SyncAliyunRouterInterfaceFromRemote(uuid string, params param.SyncAliyunRouterInterfaceFromRemoteParam) (*view.SyncAliyunRouterInterfaceFromRemoteEventView, error) {
	resp := view.SyncAliyunRouterInterfaceFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

