// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunRouterInterfaceRemote creates AliyunRouterInterfaceRemote
func (cli *ZSClient) CreateAliyunRouterInterfaceRemote(params param.CreateAliyunRouterInterfaceRemoteParam) (*view.CreateAliyunRouterInterfaceRemoteEventView, error) {
	resp := view.CreateAliyunRouterInterfaceRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/router-interface", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
