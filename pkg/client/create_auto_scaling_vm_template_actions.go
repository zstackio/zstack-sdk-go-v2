// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAutoScalingVmTemplate creates AutoScalingVmTemplate
func (cli *ZSClient) CreateAutoScalingVmTemplate(params param.CreateAutoScalingVmTemplateParam) (*view.CreateAutoScalingTemplateEventView, error) {
	resp := view.CreateAutoScalingTemplateEventView{}
	if err := cli.Post("v1/autoscaling/vmtemplate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
