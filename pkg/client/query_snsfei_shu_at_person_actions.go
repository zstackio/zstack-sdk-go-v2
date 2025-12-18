// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSFeiShuAtPerson queries SNSFeiShuAtPerson list
func (cli *ZSClient) QuerySNSFeiShuAtPerson(params param.QueryParam) ([]view.SNSFeiShuAtPersonInventoryView, error) {
	var resp []view.SNSFeiShuAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/feishu/at-persons", &params, &resp)
}
