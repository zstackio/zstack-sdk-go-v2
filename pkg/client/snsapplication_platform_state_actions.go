// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSNSApplicationPlatformState 操作SNSApplicationPlatformState
func (cli *ZSClient) ChangeSNSApplicationPlatformState(uuid string, params param.ChangeSNSApplicationPlatformStateParam) (*view.ChangeSNSApplicationPlatformStateEventView, error) {
	resp := view.ChangeSNSApplicationPlatformStateEventView{}
	if err := cli.Put("v1/sns/application-platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

