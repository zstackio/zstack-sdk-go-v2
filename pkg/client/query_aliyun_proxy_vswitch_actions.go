// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunProxyVSwitch queries AliyunProxyVSwitch list
func (cli *ZSClient) QueryAliyunProxyVSwitch(params *param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, error) {
	var resp []view.AliyunProxyVSwitchInventoryView
	return resp, cli.List("v1/aliyun-proxy/vpcs/vswitches", params, &resp)
}
