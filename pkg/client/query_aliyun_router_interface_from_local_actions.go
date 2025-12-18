// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunRouterInterfaceFromLocal queries AliyunRouterInterfaceFromLocal list
func (cli *ZSClient) QueryAliyunRouterInterfaceFromLocal(params param.QueryParam) ([]view.AliyunRouterInterfaceInventoryView, error) {
	var resp []view.AliyunRouterInterfaceInventoryView
	return resp, cli.List("v1/hybrid/aliyun/router-interface", &params, &resp)
}
