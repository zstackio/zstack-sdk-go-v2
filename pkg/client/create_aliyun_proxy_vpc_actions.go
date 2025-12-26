// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunProxyVpc creates AliyunProxyVpc
func (cli *ZSClient) CreateAliyunProxyVpc(params param.CreateAliyunProxyVpcParam) (*view.CreateAliyunProxyVpcEventView, error) {
	resp := view.CreateAliyunProxyVpcEventView{}
	if err := cli.Post("v1/aliyun-proxy/vpcs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
