// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunProxyVpc queries AliyunProxyVpc list
func (cli *ZSClient) QueryAliyunProxyVpc(params *param.QueryParam) ([]view.AliyunProxyVpcInventoryView, error) {
	var resp []view.AliyunProxyVpcInventoryView
	return resp, cli.List("v1/aliyun-proxy/vpcs", params, &resp)
}
