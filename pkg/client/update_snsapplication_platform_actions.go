// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSApplicationPlatform updates SNSApplicationPlatform
func (cli *ZSClient) UpdateSNSApplicationPlatform(uuid string, params param.UpdateSNSApplicationPlatformParam) (*view.UpdateSNSApplicationPlatformEventView, error) {
	resp := view.UpdateSNSApplicationPlatformEventView{}
	if err := cli.Put("v1/sns/application-platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
