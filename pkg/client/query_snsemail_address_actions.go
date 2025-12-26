// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSEmailAddress queries SNSEmailAddress list
func (cli *ZSClient) QuerySNSEmailAddress(params *param.QueryParam) ([]view.SNSEmailAddressInventoryView, error) {
	var resp []view.SNSEmailAddressInventoryView
	return resp, cli.List("v1/sns/application-endpoints/emails/email-addresses", params, &resp)
}
