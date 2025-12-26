// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSNSTopicState changes SNSTopicState
func (cli *ZSClient) ChangeSNSTopicState(uuid string, params param.ChangeSNSTopicStateParam) (*view.ChangeSNSTopicStateEventView, error) {
	resp := view.ChangeSNSTopicStateEventView{}
	if err := cli.Put("v1/zwatch/topics/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
