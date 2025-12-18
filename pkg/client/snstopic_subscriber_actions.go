// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSTopicSubscriber 查询SNSTopicSubscriber列表
func (cli *ZSClient) QuerySNSTopicSubscriber(params param.QueryParam) ([]view.QuerySNSTopicSubscriberView, error) {
	var resp []view.QuerySNSTopicSubscriberView
	return resp, cli.List("v1/sns/topics/subscribers", &params, &resp)
}

