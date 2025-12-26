// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSTopicSubscriber queries SNSTopicSubscriber list
func (cli *ZSClient) QuerySNSTopicSubscriber(params *param.QueryParam) ([]view.SNSSubscriberInventoryView, error) {
	var resp []view.SNSSubscriberInventoryView
	return resp, cli.List("v1/sns/topics/subscribers", params, &resp)
}
