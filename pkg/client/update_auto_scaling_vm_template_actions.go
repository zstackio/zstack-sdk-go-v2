// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAutoScalingVmTemplate updates AutoScalingVmTemplate
func (cli *ZSClient) UpdateAutoScalingVmTemplate(uuid string, params param.UpdateAutoScalingVmTemplateParam) (*view.UpdateAutoScalingTemplateEventView, error) {
	resp := view.UpdateAutoScalingTemplateEventView{}
	if err := cli.Put("v1/autoscaling/vmtemplate/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
