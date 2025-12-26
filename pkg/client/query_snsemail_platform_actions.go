// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSEmailPlatform queries SNSEmailPlatform list
func (cli *ZSClient) QuerySNSEmailPlatform(params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms/email", params, &resp)
}
