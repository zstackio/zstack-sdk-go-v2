// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHuaweiIMasterVRouter queries HuaweiIMasterVRouter list
func (cli *ZSClient) QueryHuaweiIMasterVRouter(params param.QueryParam) ([]view.HuaweiIMasterVRouterInventoryView, error) {
	var resp []view.HuaweiIMasterVRouterInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vrouters", &params, &resp)
}
