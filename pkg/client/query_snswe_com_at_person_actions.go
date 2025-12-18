// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSWeComAtPerson queries SNSWeComAtPerson list
func (cli *ZSClient) QuerySNSWeComAtPerson(params param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, error) {
	var resp []view.SNSWeComAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com/at-persons", &params, &resp)
}
