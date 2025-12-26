// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAtPersonOfAtDingTalkEndpoint updates AtPersonOfAtDingTalkEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtDingTalkEndpoint(uuid string, params param.UpdateAtPersonOfAtDingTalkEndpointParam) (*view.UpdateAtPersonOfDingTalkEndpointEventView, error) {
	resp := view.UpdateAtPersonOfDingTalkEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/ding-talk/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
