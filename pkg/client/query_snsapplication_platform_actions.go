// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSApplicationPlatform queries SNSApplicationPlatform list
func (cli *ZSClient) QuerySNSApplicationPlatform(params *param.QueryParam) ([]view.SNSApplicationPlatformInventoryView, error) {
	var resp []view.SNSApplicationPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms", params, &resp)
}
