// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunProxyVpc updates AliyunProxyVpc
func (cli *ZSClient) UpdateAliyunProxyVpc(uuid string, params param.UpdateAliyunProxyVpcParam) (*view.UpdateAliyunProxyVpcEventView, error) {
	resp := view.UpdateAliyunProxyVpcEventView{}
	if err := cli.Put("v1/aliyun-proxy/vpcs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
