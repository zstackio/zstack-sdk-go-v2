// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSTopic queries SNSTopic list
func (cli *ZSClient) QuerySNSTopic(params *param.QueryParam) ([]view.SNSTopicInventoryView, error) {
	var resp []view.SNSTopicInventoryView
	return resp, cli.List("v1/sns/topics", params, &resp)
}
