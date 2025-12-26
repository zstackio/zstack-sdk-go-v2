// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAutoScalingVmTemplate updates AutoScalingVmTemplate
func (cli *ZSClient) UpdateAutoScalingVmTemplate(uuid string, params param.UpdateAutoScalingVmTemplateParam) (*view.UpdateAutoScalingTemplateEventView, error) {
	resp := view.UpdateAutoScalingTemplateEventView{}
	if err := cli.Put("v1/autoscaling/vmtemplate/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
