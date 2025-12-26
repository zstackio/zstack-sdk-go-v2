// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SubscribeSNSTopic operates on SubscribeSNSTopic
func (cli *ZSClient) SubscribeSNSTopic(params param.SubscribeSNSTopicParam) (*view.SubscribeSNSTopicEventView, error) {
	resp := view.SubscribeSNSTopicEventView{}
	if err := cli.Post("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
