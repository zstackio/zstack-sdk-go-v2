// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSNSTopicState changes SNSTopicState
func (cli *ZSClient) ChangeSNSTopicState(uuid string, params param.ChangeSNSTopicStateParam) (*view.ChangeSNSTopicStateEventView, error) {
	resp := view.ChangeSNSTopicStateEventView{}
	if err := cli.Put("v1/zwatch/topics/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
