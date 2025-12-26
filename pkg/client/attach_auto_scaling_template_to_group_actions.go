// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachAutoScalingTemplateToGroup operates on AutoScalingTemplateToGroup
func (cli *ZSClient) AttachAutoScalingTemplateToGroup(params param.AttachAutoScalingTemplateToGroupParam) (*view.AttachAutoScalingTemplateToGroupEventView, error) {
	resp := view.AttachAutoScalingTemplateToGroupEventView{}
	if err := cli.Post("v1/autoscaling/template/{uuid}/groups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
