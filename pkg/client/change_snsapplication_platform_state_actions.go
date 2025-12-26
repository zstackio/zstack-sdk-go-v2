// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSNSApplicationPlatformState changes SNSApplicationPlatformState
func (cli *ZSClient) ChangeSNSApplicationPlatformState(uuid string, params param.ChangeSNSApplicationPlatformStateParam) (*view.ChangeSNSApplicationPlatformStateEventView, error) {
	resp := view.ChangeSNSApplicationPlatformStateEventView{}
	if err := cli.Put("v1/sns/application-platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
