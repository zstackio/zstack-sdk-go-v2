// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSTopic creates SNSTopic
func (cli *ZSClient) CreateSNSTopic(params param.CreateSNSTopicParam) (*view.CreateSNSTopicEventView, error) {
	resp := view.CreateSNSTopicEventView{}
	if err := cli.Post("v1/sns/topics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
