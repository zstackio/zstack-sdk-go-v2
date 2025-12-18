// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAtPersonOfAtDingTalkEndpoint 更新AtPersonOfAtDingTalkEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtDingTalkEndpoint(uuid string, params param.UpdateAtPersonOfAtDingTalkEndpointParam) (*view.UpdateAtPersonOfDingTalkEndpointEventView, error) {
	resp := view.UpdateAtPersonOfDingTalkEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/ding-talk/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

