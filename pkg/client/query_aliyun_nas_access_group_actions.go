// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunNasAccessGroup queries AliyunNasAccessGroup list
func (cli *ZSClient) QueryAliyunNasAccessGroup(params *param.QueryParam) ([]view.AliyunNasAccessGroupInventoryView, error) {
	var resp []view.AliyunNasAccessGroupInventoryView
	return resp, cli.List("v1/nas/aliyun/access", params, &resp)
}
