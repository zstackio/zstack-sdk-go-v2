// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunRouterInterfaceRemote creates AliyunRouterInterfaceRemote
func (cli *ZSClient) CreateAliyunRouterInterfaceRemote(params param.CreateAliyunRouterInterfaceRemoteParam) (*view.CreateAliyunRouterInterfaceRemoteEventView, error) {
	resp := view.CreateAliyunRouterInterfaceRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/router-interface", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
