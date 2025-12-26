// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunRouteEntryFromLocal queries AliyunRouteEntryFromLocal list
func (cli *ZSClient) QueryAliyunRouteEntryFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouteEntryInventoryView, error) {
	var resp []view.VpcVirtualRouteEntryInventoryView
	return resp, cli.List("v1/hybrid/aliyun/route-entry", params, &resp)
}
