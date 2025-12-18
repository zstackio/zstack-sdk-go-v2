// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunRouteInterfaceRemote updates AliyunRouteInterfaceRemote
func (cli *ZSClient) UpdateAliyunRouteInterfaceRemote(uuid string, params param.UpdateAliyunRouteInterfaceRemoteParam) (*view.UpdateAliyunRouteInterfaceRemoteEventView, error) {
	resp := view.UpdateAliyunRouteInterfaceRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
