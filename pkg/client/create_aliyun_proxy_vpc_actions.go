// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunProxyVpc creates AliyunProxyVpc
func (cli *ZSClient) CreateAliyunProxyVpc(params param.CreateAliyunProxyVpcParam) (*view.CreateAliyunProxyVpcEventView, error) {
	resp := view.CreateAliyunProxyVpcEventView{}
	if err := cli.Post("v1/aliyun-proxy/vpcs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
