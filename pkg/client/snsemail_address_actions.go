// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSEmailAddress 查询SNSEmailAddress列表
func (cli *ZSClient) QuerySNSEmailAddress(params param.QueryParam) ([]view.QuerySNSEmailAddressView, error) {
	var resp []view.QuerySNSEmailAddressView
	return resp, cli.List("v1/sns/application-endpoints/emails/email-addresses", &params, &resp)
}

