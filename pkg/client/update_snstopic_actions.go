// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSTopic updates SNSTopic
func (cli *ZSClient) UpdateSNSTopic(uuid string, params param.UpdateSNSTopicParam) (*view.UpdateSNSTopicEventView, error) {
	resp := view.UpdateSNSTopicEventView{}
	if err := cli.Put("v1/sns/topics/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
