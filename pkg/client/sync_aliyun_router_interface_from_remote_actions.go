// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncAliyunRouterInterfaceFromRemote operates on SyncAliyunRouterInterfaceFromRemote
func (cli *ZSClient) SyncAliyunRouterInterfaceFromRemote(uuid string, params param.SyncAliyunRouterInterfaceFromRemoteParam) (*view.SyncAliyunRouterInterfaceFromRemoteEventView, error) {
	resp := view.SyncAliyunRouterInterfaceFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
