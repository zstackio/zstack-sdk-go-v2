// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSNSTopic updates SNSTopic
func (cli *ZSClient) UpdateSNSTopic(uuid string, params param.UpdateSNSTopicParam) (*view.UpdateSNSTopicEventView, error) {
	resp := view.UpdateSNSTopicEventView{}
	if err := cli.Put("v1/sns/topics/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
