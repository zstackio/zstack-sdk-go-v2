// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSFeiShuAtPerson queries SNSFeiShuAtPerson list
func (cli *ZSClient) QuerySNSFeiShuAtPerson(params *param.QueryParam) ([]view.SNSFeiShuAtPersonInventoryView, error) {
	var resp []view.SNSFeiShuAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/feishu/at-persons", params, &resp)
}
