// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachAutoScalingTemplateToGroup 操作AutoScalingTemplateToGroup
func (cli *ZSClient) AttachAutoScalingTemplateToGroup(params param.AttachAutoScalingTemplateToGroupParam) (*view.AttachAutoScalingTemplateToGroupEventView, error) {
	resp := view.AttachAutoScalingTemplateToGroupEventView{}
	if err := cli.Post("v1/autoscaling/template/{uuid}/groups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

