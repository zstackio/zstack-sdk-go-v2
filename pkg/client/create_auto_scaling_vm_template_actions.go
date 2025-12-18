// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAutoScalingVmTemplate creates AutoScalingVmTemplate
func (cli *ZSClient) CreateAutoScalingVmTemplate(params param.CreateAutoScalingVmTemplateParam) (*view.CreateAutoScalingTemplateEventView, error) {
	resp := view.CreateAutoScalingTemplateEventView{}
	if err := cli.Post("v1/autoscaling/vmtemplate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
