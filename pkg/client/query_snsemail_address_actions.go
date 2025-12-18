// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSEmailAddress queries SNSEmailAddress list
func (cli *ZSClient) QuerySNSEmailAddress(params param.QueryParam) ([]view.SNSEmailAddressInventoryView, error) {
	var resp []view.SNSEmailAddressInventoryView
	return resp, cli.List("v1/sns/application-endpoints/emails/email-addresses", &params, &resp)
}
