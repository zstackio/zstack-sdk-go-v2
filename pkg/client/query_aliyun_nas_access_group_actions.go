// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunNasAccessGroup queries AliyunNasAccessGroup list
func (cli *ZSClient) QueryAliyunNasAccessGroup(params param.QueryParam) ([]view.AliyunNasAccessGroupInventoryView, error) {
	var resp []view.AliyunNasAccessGroupInventoryView
	return resp, cli.List("v1/nas/aliyun/access", &params, &resp)
}
