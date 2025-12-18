// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSTopic creates SNSTopic
func (cli *ZSClient) CreateSNSTopic(params param.CreateSNSTopicParam) (*view.CreateSNSTopicEventView, error) {
	resp := view.CreateSNSTopicEventView{}
	if err := cli.Post("v1/sns/topics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
