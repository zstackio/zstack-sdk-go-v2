// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunRouteInterfaceRemote updates AliyunRouteInterfaceRemote
func (cli *ZSClient) UpdateAliyunRouteInterfaceRemote(uuid string, params param.UpdateAliyunRouteInterfaceRemoteParam) (*view.UpdateAliyunRouteInterfaceRemoteEventView, error) {
	resp := view.UpdateAliyunRouteInterfaceRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
