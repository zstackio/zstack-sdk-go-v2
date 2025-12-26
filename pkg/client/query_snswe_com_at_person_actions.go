// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSWeComAtPerson queries SNSWeComAtPerson list
func (cli *ZSClient) QuerySNSWeComAtPerson(params *param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, error) {
	var resp []view.SNSWeComAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com/at-persons", params, &resp)
}
