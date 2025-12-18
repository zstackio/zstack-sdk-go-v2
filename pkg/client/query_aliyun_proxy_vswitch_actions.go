// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunProxyVSwitch queries AliyunProxyVSwitch list
func (cli *ZSClient) QueryAliyunProxyVSwitch(params param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, error) {
	var resp []view.AliyunProxyVSwitchInventoryView
	return resp, cli.List("v1/aliyun-proxy/vpcs/vswitches", &params, &resp)
}
